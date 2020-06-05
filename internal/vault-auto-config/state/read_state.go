package state

// A function that reads a state from the client and outputs it into the given state
type Reader func(Client, *Node) error

// A function that reads a state for a given name from the client and outputs it into the given state
type NamedReader func(Client, string, *Node) error

func ReadState(client Client) (*ConfigState, error) {
	config := NewConfigState()
	err := ReadV1State(client, config)
	return config, err
}

// Reads state using multiple state reader functions, returning an error on the first that fails
func ReadStates(client Client, node *Node, readers ...Reader) error {
	for _, reader := range readers {
		if err := reader(client, node); err != nil {
			return err
		}
	}

	return nil
}

// Reads state using multiple named state reader functions, returning an error on the first that fails
func ReadNamedStates(client Client, node *Node, name string, readers ...NamedReader) error {
	for _, reader := range readers {
		if err := reader(client, name, node); err != nil {
			return err
		}
	}

	return nil
}

// Reads the root state
func ReadV1State(client Client, state *ConfigState) error {
	node := NewNode()
	state.Children["v1"] = node
	return ReadStates(
		client,
		node,
		ReadSysState,
		ReadAuthState,
	)
}
