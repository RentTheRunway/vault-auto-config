package state

// Reads policy states
func ReadSysPolicyState(client Client, node *Node) error {
	node = node.AddNode("policy")
	policies, err := client.List("sys/policy")

	if err != nil {
		return err
	}

	for _, policy := range policies {
		policyNode := node.AddNode(policy.name)
		policyNode.Config = policy.value
	}

	return nil
}
