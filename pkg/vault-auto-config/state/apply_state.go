package state

import (
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/client"
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/config"
	v1 "github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/state/v1"
)

// Performs a diff of current and new state and applies configuration changes to vault
func ApplyState(state *config.State, client client.Client) error {
	return v1.ApplyV1State(state, client)
}
