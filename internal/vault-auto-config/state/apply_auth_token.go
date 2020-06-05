package state

// Applies state for auth backends of type "token"
func ApplyAuthTokenState(node *Node, name string, client Client) error {
	return ApplyAuthRolesState(node, name, client)
}
