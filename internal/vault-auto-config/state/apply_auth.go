package state

import (
	"errors"
	"fmt"
)

func ApplyAuthState(node *Node, client Client) error {
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
		if !ok {
			return errors.New(fmt.Sprintf("invalid state, missing required auth mount %s", name))
		}

		kind, err := GetString(auth.Config, "type")
		if err != nil {
			return err
		}

		var writer func(*Node, string, Client) error
		switch kind {
		case "kubernetes":
			writer = ApplyAuthKubernetesState
		case "okta":
			writer = ApplyAuthOktaState
		case "token":
			writer = ApplyAuthTokenState
		}

		if writer == nil {
			return errors.New(fmt.Sprintf("unable to write state for unsupported auth type '%s'", kind))
		}

		if err := writer(value, name, client); err != nil {
			return err
		}
	}

	return nil
}
