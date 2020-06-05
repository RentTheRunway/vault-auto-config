package state

// Reads config state for an auth backend
func ReadAuthConfigState(client Client, name string, node *Node) error {
	node = node.AddNode("config")

	config, err := client.Read("auth/%s/config", name)
	if err != nil {
		return err
	}

	node.Config = config

	return nil
}

// Reads a generic auth resource state for an auth backend (e.g. groups, roles, etc.)
func ReadAuthResourceState(client Client, name string, listResource string, resource string, node *Node) error {
	node = node.AddNode(resource)

	resources, err := client.List("auth/%s/%s", name, listResource)

	if err != nil {
		return err
	}

	for _, resource := range resources {
		resourceNode := node.AddNode(resource.name)
		resourceNode.Config = resource.value
	}

	return nil
}

// Reads group states for an auth backend
func ReadAuthGroupsState(client Client, name string, node *Node) error {
	return ReadAuthResourceState(client, name, "groups", "groups", node)
}

// Reads user states for an auth backend
func ReadAuthUsersState(client Client, name string, node *Node) error {
	return ReadAuthResourceState(client, name, "users", "users", node)
}

// Reads role states for an auth backend
func ReadAuthRolesState(client Client, name string, node *Node) error {
	return ReadAuthResourceState(client, name, "roles", "roles", node)
}

// Reads role states for an auth backend, but with the singular name "role"
func ReadAuthRoleState(client Client, name string, node *Node) error {
	return ReadAuthResourceState(client, name, "role", "role", node)
}
