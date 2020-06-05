package state

func ApplyAuthKubernetesState(node *Node, name string, client Client) error {
	return ApplyNamedStates(
		node,
		client,
		name,
		ApplyAuthRoleState,
		ApplyAuthConfigState,
	)
}
