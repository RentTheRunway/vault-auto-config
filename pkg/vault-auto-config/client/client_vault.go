package client

import (
	"fmt"
	"strings"

	yaml2 "github.com/goccy/go-yaml"
	"github.com/hashicorp/vault/api"
)

// A client for reading and writing vault configuration state using the vault api
type VaultClient struct {
	client *api.Client
}

var readOnlyPaths = map[string]bool{
	"sys/auth/token":     true,
	"sys/policy/default": true,
	"sys/policy/root":    true,
}

var writeOnlyEndpoints = map[string]bool{
	"secret-id": true,
}

// Creates a new VaultClient
func NewVaultClient(url string, token string) (*VaultClient, error) {
	client, err := api.NewClient(&api.Config{Address: url})

	if err != nil {
		return nil, err
	}

	client.SetToken(token)
	return &VaultClient{client: client}, nil
}

// Lists a resource path
func (c *VaultClient) List(path string, args ...interface{}) (Entries, error) {
	path = fmt.Sprintf(path, args...)

	if strings.HasPrefix(path, "sys/") {
		return c.listSys(path)
	} else {
		return c.list(path)
	}
}

// Auth is listed differently that everything else
func (c *VaultClient) listSys(path string) (Entries, error) {
	switch path {
	case "sys/auth":
		return c.listSysAuth()
	case "sys/policy":
		return c.listSysPolicy()
	default:
		return nil, fmt.Errorf("unsupported sys path: %s", path)
	}
}

func (c *VaultClient) listSysAuth() ([]*Entry, error) {
	log.Debugf("Listing api resources v1/sys/auth")

	result, err := c.client.Sys().ListAuth()

	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	var data []*Entry = nil

	// payload contains the sub-resources themselves
	for key, value := range result {
		name := key[0 : len(key)-1]

		var mapped Payload
		bytes, err := yaml2.Marshal(value)
		if err != nil {
			return nil, err
		}
		if err := yaml2.Unmarshal(bytes, &mapped); err != nil {
			return nil, err
		}

		data = append(data, &Entry{Name: name, Value: mapped})
	}

	return data, nil
}

// For paths other than sys
func (c *VaultClient) listSysPolicy() (Entries, error) {
	log.Debugf("Listing api resources v1/sys/policy")

	policies, err := c.client.Sys().ListPolicies()

	if err != nil {
		return nil, err
	}

	var data []*Entry = nil

	for _, name := range policies {
		value, err := c.client.Sys().GetPolicy(name)
		if err != nil {
			return nil, err
		}
		policy := make(map[string]string)
		policy["policy"] = value

		data = append(data, &Entry{Name: name, Value: policy})
	}

	return data, err
}

// For paths other than sys
func (c *VaultClient) list(path string) (Entries, error) {
	log.Debugf("Listing api resources %s", path)

	result, err := c.client.Logical().List(path)

	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	keys, ok := result.Data["keys"].([]interface{})

	if !ok {
		return nil, fmt.Errorf("keys for %s of the wrong type", path)
	}

	var data []*Entry = nil

	// payload has a 'keys' field with sub-resources to fetch
	for _, key := range keys {
		value, err := c.client.Logical().Read(fmt.Sprintf("%s/%s", path, key))
		if err != nil {
			return nil, err
		}

		keyString, ok := key.(string)
		if !ok {
			return nil, fmt.Errorf("key for %s of the wrong type", path)
		}

		data = append(data, &Entry{Name: keyString, Value: value.Data})
	}

	return data, err
}

// Writes a resource
func (c *VaultClient) Write(data Payload, path string, args ...interface{}) error {
	path = fmt.Sprintf(path, args...)

	if _, ok := readOnlyPaths[path]; ok {
		return nil
	}

	log.Debugf("Writing api resource %s", path)

	mapped, ok := data.(map[string]interface{})
	if !ok {
		return fmt.Errorf("could not write data for path '%s', wrong type", path)
	}
	_, err := c.client.Logical().Write(path, mapped)
	return err
}

// Reads a resource
func (c *VaultClient) Read(path string, args ...interface{}) (Payload, error) {
	path = fmt.Sprintf(path, args...)

	endpoint := path[strings.LastIndex(path, "/")+1:]

	if _, ok := writeOnlyEndpoints[endpoint]; ok {
		return nil, nil
	}

	log.Debugf("Reading api resource %s", path)

	data, err := c.client.Logical().Read(path)

	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, nil
	}

	return data.Data, nil
}

// Deletes a resource
func (c *VaultClient) Delete(path string, args ...interface{}) error {
	path = fmt.Sprintf(path, args...)

	if _, ok := readOnlyPaths[path]; ok {
		return nil
	}

	log.Debugf("Deleting api resource %s", path)

	_, err := c.client.Logical().Delete(path)
	return err
}
