package auth

import (
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/client"
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/config"
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/util"
)

// Reads state for auth backends of type "kubernetes"
func ReadAuthKubernetesState(client client.Client, name string, node *config.Node) error {
	return util.ReadNamedStates(
		client,
		node,
		name,
		ReadAuthRoleState,
		ReadAuthConfigState,
	)
}
