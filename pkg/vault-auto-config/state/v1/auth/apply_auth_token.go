package auth

import (
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/client"
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/config"
)

// Applies state for auth backends of type "token"
func ApplyAuthTokenState(node *config.Node, name string, client client.Client) error {
	return ApplyAuthRolesState(node, name, client)
}
