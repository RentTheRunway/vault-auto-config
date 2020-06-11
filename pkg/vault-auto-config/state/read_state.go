package state

import (
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/client"
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/config"
	v1 "github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/state/v1"
)

func ReadState(client client.Client) (*config.State, error) {
	state := config.NewState()
	err := v1.ReadV1State(client, state)
	return state, err
}
