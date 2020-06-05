package state

func ApplyAuthTokenState(node *Node, name string, client Client) error {
	return ApplyAuthRolesState(node, name, client)
}
