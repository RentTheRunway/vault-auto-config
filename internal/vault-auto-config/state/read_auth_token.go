package state

func ReadAuthTokenState(client Client, name string, node *Node) error {
	return ReadAuthRolesState(client, name, node)
}
