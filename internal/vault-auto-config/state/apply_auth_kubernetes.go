package state

// Applies state for auth backends of type "kubernetes"
func ApplyAuthKubernetesState(node *Node, name string, client Client) error {
	return ApplyNamedStates(
		node,
		client,
		name,
		ApplyAuthRoleState,
		ApplyAuthConfigState,
	)
}
