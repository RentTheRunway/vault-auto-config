package sys

import (
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/client"
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/config"
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/state/v1/sys/auth"
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/state/v1/sys/policy"
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/util"
)

// Applies sys state
func ApplySysState(node *config.Node, client client.Client) error {
	node, ok := node.Children["sys"]
	if !ok {
		return nil
	}
	return util.ApplyStates(
		node,
		client,
		auth.ApplySysAuthState,
		policy.ApplySysPolicyState,
	)
}
