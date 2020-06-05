package state

// Applies auth backend states
func ApplySysAuthState(node *Node, client Client) error {
	node, ok := node.Children["auth"]
	if !ok {
		return nil
	}

	existing, err := client.List("sys/auth")
	if err != nil {
		return err
	}

	// prune configs
	for _, entry := range existing {
		if _, ok := node.Children[entry.name]; !ok {
			if err := client.Delete("sys/auth/%s", entry.name); err != nil {
				return err
			}
		}
	}

	// apply
	for name, value := range node.Children {
		if existing.Exists(name) {
			if err := client.Delete("sys/auth/%s", name); err != nil {
				return err
			}
		}

		if value.Config != nil {
			if err := client.Write(value.Config, "sys/auth/%s", name); err != nil {
				return err
			}
		}
	}

	return nil
}
