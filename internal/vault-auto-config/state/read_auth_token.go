package state

// Reads state for auth backends of type "token"
func ReadAuthTokenState(client Client, name string, node *Node) error {
	return ReadAuthRolesState(client, name, node)
}
