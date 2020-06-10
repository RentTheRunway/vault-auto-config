package state

import (
	"fmt"
)

// Applies auth state
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
		if !ok && name != "token" {
			return fmt.Errorf("invalid state, missing required auth mount %s", name)
		}

		var kind string
		var err error

		if name == "token" {
			kind = "token"
		} else {
			kind, err = GetString(auth.Config, "type")
			if err != nil {
				return err
			}
		}

		var writer NamedWriter
		switch kind {
		case "kubernetes":
			writer = ApplyAuthKubernetesState
		case "okta":
			writer = ApplyAuthOktaState
		case "token":
			writer = ApplyAuthTokenState
		}

		if writer == nil {
			return fmt.Errorf("unable to write state for unsupported auth type '%s'", kind)
		}

		if err := writer(value, name, client); err != nil {
			return err
		}
	}

	return nil
}
