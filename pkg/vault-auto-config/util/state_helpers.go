package util

import (
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/client"
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/config"
)

// A function that reads a state from the client and outputs it into the given state
type Reader func(client.Client, *config.Node) error

// A function that reads a state for a given name from the client and outputs it into the given state
type NamedReader func(client.Client, string, *config.Node) error

// A function that takes a given state and writes it to the target client
type Writer func(*config.Node, client.Client) error

// A function that takes a given state and name and writes it to the target client
type NamedWriter func(*config.Node, string, client.Client) error

// Reads state using multiple state reader functions, returning an error on the first that fails
func ReadStates(client client.Client, node *config.Node, readers ...Reader) error {
	for _, reader := range readers {
		if err := reader(client, node); err != nil {
			return err
		}
	}

	return nil
}

// Reads state using multiple named state reader functions, returning an error on the first that fails
func ReadNamedStates(client client.Client, node *config.Node, name string, readers ...NamedReader) error {
	for _, reader := range readers {
		if err := reader(client, name, node); err != nil {
			return err
		}
	}

	return nil
}

// Apply state using multiple state writer functions, returning an error on the first that fails
func ApplyStates(node *config.Node, client client.Client, writers ...Writer) error {
	for _, writer := range writers {
		if err := writer(node, client); err != nil {
			return err
		}
	}

	return nil
}

// Apply a state using multiple named state writer functions, returning an err on the first one that fails
func ApplyNamedStates(node *config.Node, client client.Client, name string, writers ...NamedWriter) error {
	for _, writer := range writers {
		if err := writer(node, name, client); err != nil {
			return err
		}
	}

	return nil
}
