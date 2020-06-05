package state

func ReadAuthKubernetesState(client Client, name string, node *Node) error {
	return ReadNamedStates(
		client,
		node,
		name,
		ReadAuthRoleState,
		ReadAuthConfigState,
	)
}
