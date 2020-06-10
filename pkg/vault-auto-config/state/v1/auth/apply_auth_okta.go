package auth

import (
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/client"
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/config"
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/util"
)

// Applies state for auth backends of type "okta"
func ApplyAuthOktaState(node *config.Node, name string, client client.Client) error {
	if err := util.ApplyNamedStates(
		node,
		client,
		name,
		ApplyAuthUsersState,
		ApplyAuthGroupsState,
	); err != nil {
		return err
	}

	return ApplyAuthConfigState(node, name, client)
}
