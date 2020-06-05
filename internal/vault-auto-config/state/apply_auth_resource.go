package state

// Applies config state for an auth backend
func ApplyAuthConfigState(node *Node, name string, client Client) error {
	node, ok := node.Children["config"]
	if !ok || node.Config == nil {
		return nil
	}

	return client.Write(node.Config, "auth/%s/config", name)
}

// Applies a generic auth resource state for an auth backend (e.g. groups, roles, etc.)
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

// Applies group states for an auth backend
func ApplyAuthGroupsState(node *Node, name string, client Client) error {
	return ApplyAuthResourceState(node, name, "groups", client)
}

// Applies user states for an auth backend
func ApplyAuthUsersState(node *Node, name string, client Client) error {
	return ApplyAuthResourceState(node, name, "users", client)
}

// Applies role states for an auth backend
func ApplyAuthRolesState(node *Node, name string, client Client) error {
	return ApplyAuthResourceState(node, name, "roles", client)
}

// Applies role states for an auth backend, but with the singular name "role"
func ApplyAuthRoleState(node *Node, name string, client Client) error {
	return ApplyAuthResourceState(node, name, "role", client)
}
