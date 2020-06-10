package sys

import (
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/client"
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/config"
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/state/v1/sys/auth"
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/state/v1/sys/policy"
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/util"
)

// Reads sys state
func ReadSysState(client client.Client, node *config.Node) error {
	node = node.AddNode("sys")

	return util.ReadStates(
		client,
		node,
		auth.ReadSysAuthState,
		policy.ReadSysPolicyState,
	)
}
