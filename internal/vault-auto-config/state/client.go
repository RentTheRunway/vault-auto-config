package state

import (
	"errors"
	"fmt"
)

type Client interface {
	List(path string, args ...interface{}) (Entries, error)
	Write(data Payload, path string, args ...interface{}) error
	Read(path string, args ...interface{}) (Payload, error)
	Delete(path string, args ...interface{}) error
}

type Payload interface{}
type Entry struct {
	name string
	value Payload
}

type Entries []*Entry

func (e Entries) Exists(name string) bool {
	for _, entry := range e {
		if entry.name == name {
			return true
		}
	}

	return false
}

func GetString(p Payload, name string) (string, error) {
	m, ok := p.(map[string]interface{})
	if !ok {
		return "", errors.New(fmt.Sprintf("could not get string '%s', payload is wrong type", name))
	}

	value, ok := m[name]
	if !ok {
		return "", nil
	}

	casted, ok := value.(string)
	if !ok {
		return "", errors.New(fmt.Sprintf("could not get string '%s', wrong type", name))
	}

	return casted, nil
}
