package state

import (
	"fmt"
)

// A generic client that is capable for reading and writing vault configuration state
type Client interface {
	List(path string, args ...interface{}) (Entries, error)
	Write(data Payload, path string, args ...interface{}) error
	Read(path string, args ...interface{}) (Payload, error)
	Delete(path string, args ...interface{}) error
}

// A config state payload
type Payload interface{}

// An named entry for a config state payload
type Entry struct {
	name  string
	value Payload
}

// An array of entries
type Entries []*Entry

// Utility method to check if an Entries contains a given name
func (e Entries) Exists(name string) bool {
	for _, entry := range e {
		if entry.name == name {
			return true
		}
	}

	return false
}

// Utility method to safely extract a field from an arbitrary payload
func GetString(p Payload, name string) (string, error) {
	m, ok := p.(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("could not get string '%s', payload is wrong type", name)
	}

	value, ok := m[name]
	if !ok {
		return "", nil
	}

	casted, ok := value.(string)
	if !ok {
		return "", fmt.Errorf("could not get string '%s', wrong type", name)
	}

	return casted, nil
}
