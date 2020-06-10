package v1

import (
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/client"
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/config"
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/state/v1/auth"
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/state/v1/sys"
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/util"
)

// Reads the root state
func ReadV1State(client client.Client, configState *config.State) error {
	node := config.NewNode()
	configState.Children["v1"] = node
	return util.ReadStates(
		client,
		node,
		sys.ReadSysState,
		auth.ReadAuthState,
	)
}
