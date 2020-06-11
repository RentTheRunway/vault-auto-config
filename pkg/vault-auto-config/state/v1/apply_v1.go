package v1

import (
	"errors"
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/client"
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/config"
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/state/v1/auth"
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/state/v1/sys"
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/util"
)

// Applies root state
func ApplyV1State(configState *config.State, client client.Client) error {
	node, ok := configState.Children["v1"]
	if !ok {
		return errors.New("invalid state, missing required key v1")
	}

	return util.ApplyStates(
		node,
		client,
		sys.ApplySysState,
		auth.ApplyAuthState,
	)
}
