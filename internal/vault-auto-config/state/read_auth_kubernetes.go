package state

// Reads state for auth backends of type "kubernetes"
func ReadAuthKubernetesState(client Client, name string, node *Node) error {
	return ReadNamedStates(
		client,
		node,
		name,
		ReadAuthRoleState,
		ReadAuthConfigState,
	)
}
