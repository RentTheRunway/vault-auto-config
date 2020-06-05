package state

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

type FileSystemClient struct {
	dir string
	secrets map[string]interface{}
}

type templateData struct {
	Secrets map[string]interface{}
}

func NewFileSystemClient(dir string, secretsFile string) (*FileSystemClient, error) {
	var secrets map[string]interface{}

	// decrypt secret values using sops
	if secretsFile != "" {
		decrypted, err := decrypt.File(secretsFile, "yaml")
		if err != nil {
			return nil, err
		}

		if err = yaml2.Unmarshal(decrypted, &secrets); err != nil {
			return nil, err
		}
	}

	return &FileSystemClient{dir: dir, secrets: secrets}, nil
}

func (c *FileSystemClient) List(path string, args ...interface{}) (Entries, error) {
	var entries []*Entry

	err := filepath.Walk(c.resolvePath(path, args), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
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
		entries = append(entries, &Entry{name: name, value: data})
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

func (c *FileSystemClient) Write(data Payload, path string, args ...interface{}) error {
	path = fmt.Sprintf("%s.yaml", c.resolvePath(path, args))
	yaml, err := yaml2.Marshal(data)
	if err != nil {
		return err
	}

	_ = os.MkdirAll(filepath.Dir(path), 777)
	return ioutil.WriteFile(path, yaml, 666)
}

func (c *FileSystemClient) Read(path string, args ...interface{}) (Payload, error) {
	path = fmt.Sprintf("%s.yaml", c.resolvePath(path, args))
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	contents, err = c.processTemplate(contents)
	if err != nil {
		return nil, err
	}

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
	if err := tmpl.Execute(&out, templateData{ Secrets: c.secrets}); err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}

func (c *FileSystemClient) Delete(path string, args ...interface{}) error {
	path = fmt.Sprintf("%s.yaml", c.resolvePath(path, args))
	return os.Remove(path)
}

func (c *FileSystemClient) resolvePath(path string, args []interface{}) string {
	path = fmt.Sprintf(path, args...)
	path = filepath.Join(c.dir, "v1", path)
	return path
}
