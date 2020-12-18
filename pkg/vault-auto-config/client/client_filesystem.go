package client

import (
	"bytes"
	"fmt"
	yaml2 "github.com/goccy/go-yaml"
	"go.mozilla.org/sops/v3/decrypt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

// A client for reading and writing vault configuration state using a file system
type FileSystemClient struct {
	dir     string
	secrets map[string]interface{}
}

type templateData struct {
	Secrets map[string]interface{}
}

// Creates a new FileSystemClient
func NewFileSystemClient(dir string, secretsFile string, valuesFile string) (*FileSystemClient, error) {
	secrets := make(map[string]interface{})

	// add values to be templated
	if valuesFile != "" {
		log.Debugf("Loading values")
		yamlFile, err := ioutil.ReadFile(valuesFile)
		if err != nil {
			return nil, err
		}

		var values map[string]interface{}
		if err = yaml2.Unmarshal(yamlFile, &values); err != nil {
			return nil, err
		}

		for k, v := range values {
			secrets[k] = v
		}
	}

	// decrypt secret values using sops
	if secretsFile != "" {
		log.Debugf("Loading secrets")
		decrypted, err := decrypt.File(secretsFile, "yaml")
		if err != nil {
			return nil, err
		}

		var values map[string]interface{}
		if err = yaml2.Unmarshal(decrypted, &values); err != nil {
			return nil, err
		}

		for k, v := range values {
			secrets[k] = v
		}
	}

	return &FileSystemClient{dir: dir, secrets: secrets}, nil
}

// Lists config state yaml files in the path
func (c *FileSystemClient) List(path string, args ...interface{}) (Entries, error) {
	var entries []*Entry
	resolvedPath := c.resolvePath(path, args)

	log.Debugf("Listing files %s", resolvedPath)

	err := filepath.Walk(resolvedPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// don't walk recursively
		if info.IsDir() {
			if path != resolvedPath {
				return filepath.SkipDir
			}

			return nil
		}

		ext := filepath.Ext(path)
		if ext != ".yaml" {
			return nil
		}

		yaml, err := ioutil.ReadFile(path)

		if err != nil {
			return err
		}

		var data map[string]interface{}
		if err = yaml2.Unmarshal(yaml, &data); err != nil {
			return err
		}

		name := filepath.Base(path)
		name = name[0 : len(name)-len(ext)]
		entries = append(entries, &Entry{Name: name, Value: data})
		return nil
	})

	if os.IsNotExist(err) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return entries, nil
}

// Writes config state to a yaml file
func (c *FileSystemClient) Write(data Payload, path string, args ...interface{}) error {
	path = fmt.Sprintf("%s.yaml", c.resolvePath(path, args))

	log.Debugf("Writing file %s", path)

	yaml, err := yaml2.Marshal(data)
	if err != nil {
		return err
	}

	_ = os.MkdirAll(filepath.Dir(path), 0755)
	return ioutil.WriteFile(path, yaml, 0666)
}

// Reads a config state yaml file
func (c *FileSystemClient) Read(path string, args ...interface{}) (Payload, error) {
	path = fmt.Sprintf("%s.yaml", c.resolvePath(path, args))

	log.Debugf("Reading file %s", path)

	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	contents, err = c.processTemplate(contents)
	if err != nil {
		return nil, err
	}

	log.Debugf("File contents: \n%s", string(contents))

	var data Payload
	if err = yaml2.Unmarshal(contents, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func (c *FileSystemClient) processTemplate(data []byte) ([]byte, error) {
	tmpl, err := template.New("template").Parse(string(data))
	if err != nil {
		return nil, err
	}

	out := bytes.Buffer{}
	if err := tmpl.Execute(&out, templateData{Secrets: c.secrets}); err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}

// Deletes a config state yaml file
func (c *FileSystemClient) Delete(path string, args ...interface{}) error {
	path = fmt.Sprintf("%s.yaml", c.resolvePath(path, args))

	log.Debugf("Deleting file %s", path)

	return os.Remove(path)
}

func (c *FileSystemClient) resolvePath(path string, args []interface{}) string {
	path = fmt.Sprintf(path, args...)
	path = filepath.Join(c.dir, "v1", path)
	return path
}
