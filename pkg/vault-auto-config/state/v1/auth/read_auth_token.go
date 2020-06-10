package auth

import (
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/client"
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/config"
)

// Reads state for auth backends of type "token"
func ReadAuthTokenState(client client.Client, name string, node *config.Node) error {
	return ReadAuthRolesState(client, name, node)
}
