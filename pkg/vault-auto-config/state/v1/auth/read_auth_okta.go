package auth

import (
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/client"
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/config"
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/util"
)

// Reads state for auth backends of type "okta"
func ReadAuthOktaState(client client.Client, name string, node *config.Node) error {
	if err := util.ReadNamedStates(
		client,
		node,
		name,
		ReadAuthUsersState,
		ReadAuthGroupsState,
	); err != nil {
		return err
	}

	return ReadAuthConfigState(client, name, node)
}
