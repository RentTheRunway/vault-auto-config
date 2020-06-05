package state

import (
	"errors"
	"fmt"
)

func ReadAuthState(client Client, node *Node) error {
	node = node.AddNode("auth")
	auths, err := client.List("sys/auth")

	if err != nil {
		return err
	}

	for _, auth := range auths {
		kind, err :=  GetString(auth.value, "type")
		if err != nil {
			return err
		}

		authNode := node.AddNode(auth.name)

		var reader func(Client, string, *Node) error
		switch kind {
		case "kubernetes":
			reader = ReadAuthKubernetesState
		case "okta":
			reader = ReadAuthOktaState
		case "token":
			reader = ReadAuthTokenState
		}

		if reader == nil {
			return errors.New(fmt.Sprintf("unable to read state for unsupported auth type '%s'", kind))
		}

		if err := reader(client, auth.name, authNode); err != nil {
			return err
		}
	}

	return nil
}
