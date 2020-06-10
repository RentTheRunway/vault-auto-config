package policy

import (
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/client"
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/config"
)

// Reads policy states
func ReadSysPolicyState(client client.Client, node *config.Node) error {
	node = node.AddNode("policy")
	policies, err := client.List("sys/policy")

	if err != nil {
		return err
	}

	for _, policy := range policies {
		policyNode := node.AddNode(policy.Name)
		policyNode.Config = policy.Value
	}

	return nil
}
