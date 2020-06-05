package state

// Applies state for auth backends of type "okta"
func ApplyAuthOktaState(node *Node, name string, client Client) error {
	if err := ApplyNamedStates(
		node,
		client,
		name,
		ApplyAuthUsersState,
		ApplyAuthGroupsState,
	); err != nil {
		return err
	}

	return ApplyAuthConfigState(node, name, client)
}
