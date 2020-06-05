package state

import "errors"

// A function that takes a given state and writes it to the target client
type Writer func(*Node, Client)error
// A function that takes a given state and name and writes it to the target client
type NamedWriter func(*Node, string, Client)error


// Performs a diff of current and new state and applies configuration changes to vault
func ApplyState(state *ConfigState, client Client) error {
	return ApplyV1State(state, client)
}

// Apply state using multiple state writer functions, returning an error on the first that fails
func ApplyStates(node *Node, client Client, writers ...Writer) error {
	for _, writer := range writers {
		if err := writer(node, client); err != nil {
			return err
		}
	}

	return nil
}

func ApplyNamedStates(node *Node, client Client, name string, writers ...NamedWriter) error {
	for _, writer := range writers {
		if err := writer(node, name, client); err != nil {
			return err
		}
	}

	return nil
}

func ApplyV1State(state *ConfigState, client Client) error {
	node, ok := state.Children["v1"]
	if !ok {
		return errors.New("invalid state, missing required key v1")
	}

	return ApplyStates(
		node,
		client,
		ApplySysState,
		ApplyAuthState,
	)
}
