package state


func ApplyAuthConfigState(node *Node, name string, client Client) error {
	node, ok := node.Children["config"]
	if !ok || node.Config == nil {
		return nil
	}

	return client.Write(node.Config, "auth/%s/config", name)
}

func ApplyAuthResourceState(node *Node, name string, resource string, client Client) error {
	node = node.Children[resource]

	// prune
	existing, err := client.List("auth/%s/%s", name, resource)
	if err != nil {
		return err
	}

	for _, entry := range existing {
		remove := false

		if node == nil {
			remove = true
		} else if _, ok := node.Children[entry.name]; !ok {
			remove = true
		}

		if remove {
			if err := client.Delete("auth/%s/%s/%s", name, resource, entry.name); err != nil {
				return err
			}
		}
	}

	// apply
	if node != nil {
		for key, value := range node.Children {
			if value.Config == nil {
				continue
			}

			if err := client.Write(value.Config, "auth/%s/%s/%s", name, resource, key); err != nil {
				return err
			}
		}
	}

	return nil
}

func ApplyAuthGroupsState(node *Node, name string, client Client) error {
	return ApplyAuthResourceState(node, name, "groups", client)
}

func ApplyAuthUsersState(node *Node, name string, client Client) error {
	return ApplyAuthResourceState(node, name, "users", client)
}

func ApplyAuthRolesState(node *Node, name string, client Client) error {
	return ApplyAuthResourceState(node, name, "roles", client)
}

func ApplyAuthRoleState(node *Node, name string, client Client) error {
	return ApplyAuthResourceState(node, name, "role", client)
}
