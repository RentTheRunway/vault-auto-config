package auth

import (
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/client"
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/config"
)

// Applies auth backend states
func ApplySysAuthState(node *config.Node, client client.Client) error {
	node, ok := node.Children["auth"]
	if !ok {
		return nil
	}

	existing, err := client.List("sys/auth")
	if err != nil {
		return err
	}

	// prune configs
	for _, entry := range existing {
		if _, ok := node.Children[entry.Name]; !ok {
			if err := client.Delete("sys/auth/%s", entry.Name); err != nil {
				return err
			}
		}
	}

	// apply
	for name, value := range node.Children {
		if existing.Exists(name) {
			if err := client.Delete("sys/auth/%s", name); err != nil {
				return err
			}
		}

		if value.Config != nil {
			if err := client.Write(value.Config, "sys/auth/%s", name); err != nil {
				return err
			}
		}
	}

	return nil
}
