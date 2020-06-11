package auth

import (
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/client"
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/config"
)

// Reads auth backend states
func ReadSysAuthState(client client.Client, node *config.Node) error {
	node = node.AddNode("auth")
	auths, err := client.List("sys/auth")

	if err != nil {
		return err
	}

	for _, auth := range auths {
		authNode := node.AddNode(auth.Name)
		authNode.Config = auth.Value
	}

	return nil
}
