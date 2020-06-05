package state

func ReadAuthConfigState(client Client, name string, node *Node) error {
	node = node.AddNode("config")

	config, err := client.Read("auth/%s/config", name)
	if err != nil {
		return err
	}

	node.Config = config

	return nil
}

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

func ReadAuthGroupsState(client Client, name string, node *Node) error {
	return ReadAuthResourceState(client, name, "groups", "groups", node)
}

func ReadAuthUsersState(client Client, name string, node *Node) error {
	return ReadAuthResourceState(client, name, "users", "users", node)
}

func ReadAuthRolesState(client Client, name string, node *Node) error {
	return ReadAuthResourceState(client, name, "roles", "roles", node)
}

func ReadAuthRoleState(client Client, name string, node *Node) error {
	return ReadAuthResourceState(client, name, "role", "role", node)
}
