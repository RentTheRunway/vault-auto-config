package auth

import (
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/client"
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/config"
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/util"
)

// Applies state for auth backends of type "approle"
func ApplyAuthApproleState(node *config.Node, name string, client client.Client) error {
	return util.ApplyNamedStates(
		node,
		client,
		name,
		ApplyAuthRoleState,
		ApplyAuthApproleRoleIdState,
		ApplyAuthApproleSecretIdState,
	)
}
