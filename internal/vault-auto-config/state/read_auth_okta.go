package state

func ReadAuthOktaState(client Client, name string, node *Node) error {
	if err := ReadNamedStates(
		client,
		node,
		name,
		ReadAuthUsersState,
		ReadAuthGroupsState,
	); err != nil {
		return err
	}

	return ReadAuthConfigState(client, name, node)
}
