package state

func ReadSysAuthState(client Client, node *Node) error {
	node = node.AddNode("auth")
	auths, err := client.List("sys/auth")

	if err != nil {
		return err
	}

	for _, auth := range auths {
		authNode := node.AddNode(auth.name)
		authNode.Config = auth.value
	}

	return nil
}
