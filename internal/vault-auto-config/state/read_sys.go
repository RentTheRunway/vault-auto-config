package state

func ReadSysState(client Client, node *Node) error {
	node = node.AddNode("sys")

	return ReadStates(
		client,
		node,
		ReadSysAuthState,
		ReadSysPolicyState,
	)
}
