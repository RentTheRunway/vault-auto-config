package auth

import (
	"fmt"
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/client"
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/config"
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/util"
)

// Read auth state
func ReadAuthState(stateClient client.Client, node *config.Node) error {
	node = node.AddNode("auth")
	auths, err := stateClient.List("sys/auth")

	if err != nil {
		return err
	}

	for _, authConfig := range auths {
		kind, err := client.GetString(authConfig.Value, "type")
		if err != nil {
			return err
		}

		authNode := node.AddNode(authConfig.Name)

		var reader util.NamedReader
		switch kind {
		case "kubernetes":
			reader = ReadAuthKubernetesState
		case "okta":
			reader = ReadAuthOktaState
		case "approle":
			reader = ReadAuthApproleState
		case "token":
			continue
		}

		if reader == nil {
			return fmt.Errorf("unable to read state for unsupported auth type '%s'", kind)
		}

		if err := reader(stateClient, authConfig.Name, authNode); err != nil {
			return err
		}
	}

	// token is built-in and assumed to always exist
	if err := ReadAuthTokenState(stateClient, "token", node.AddNode("token")); err != nil {
		return err
	}

	return nil
}
