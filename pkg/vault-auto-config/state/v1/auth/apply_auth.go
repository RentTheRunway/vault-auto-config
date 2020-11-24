package auth

import (
	"fmt"
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/client"
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/config"
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/util"
)

// Applies auth state
func ApplyAuthState(node *config.Node, stateClient client.Client) error {
	node, ok := node.Children["auth"]
	if !ok {
		return nil
	}

	sys, ok := node.Parent.Children["sys"]
	if !ok {
		return nil
	}

	auths, ok := sys.Children["auth"]
	if !ok {
		return nil
	}

	for name, value := range node.Children {
		auth, ok := auths.Children[name]
		if !ok && name != "token" {
			return fmt.Errorf("invalid state, missing required auth mount %s", name)
		}

		var kind string
		var err error

		if name == "token" {
			kind = "token"
		} else {
			kind, err = client.GetString(auth.Config, "type")
			if err != nil {
				return err
			}
		}

		var writer util.NamedWriter
		switch kind {
		case "kubernetes":
			writer = ApplyAuthKubernetesState
		case "okta":
			writer = ApplyAuthOktaState
		case "approle":
			writer = ApplyAuthApplroleState
		case "token":
			writer = ApplyAuthTokenState
		}

		if writer == nil {
			return fmt.Errorf("unable to write state for unsupported auth type '%s'", kind)
		}

		if err := writer(value, name, stateClient); err != nil {
			return err
		}
	}

	return nil
}
