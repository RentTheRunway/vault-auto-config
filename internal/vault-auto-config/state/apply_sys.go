package state

func ApplySysState(node *Node, client Client) error {
	node, ok := node.Children["sys"]
	if !ok {
		return nil
	}
	return ApplyStates(
		node,
		client,
		ApplySysAuthState,
		ApplySysPolicyState,
	)
}
