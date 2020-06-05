package state

func ApplySysPolicyState(node* Node, client Client) error {
	node, ok := node.Children["policy"]
	if !ok {
		return nil
	}

	existing, err := client.List("sys/policy")
	if err != nil {
		return err
	}

	// prune configs
	for _, entry := range existing {
		if _, ok := node.Children[entry.name]; !ok {
			if err := client.Delete("sys/policy/%s", entry.name); err != nil {
				return err
			}
		}
	}

	// apply
	for name, value := range node.Children {
		if existing.Exists(name) {
			if err := client.Delete("sys/policy/%s", name); err != nil {
				return err
			}
		}

		if value.Config != nil {
			if err := client.Write(value.Config, "sys/policy/%s", name); err != nil {
				return err
			}
		}
	}

	return nil
}
