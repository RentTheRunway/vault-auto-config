package auth

import (
	"fmt"

	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/client"
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/config"
)

// Reads config state for an auth backend
func ReadAuthConfigState(client client.Client, name string, node *config.Node) error {
	node = node.AddNode("config")

	var err error
	if node.Config, err = client.Read("auth/%s/config", name); err != nil {
		return err
	}

	return nil
}

// Reads a generic auth resource state for an auth backend (e.g. groups, roles, etc.)
func ReadAuthResourceState(resourceClient client.Client, name string, listResource string, resource string, node *config.Node) error {
	node = node.AddNode(resource)

	resources, err := resourceClient.List("auth/%s/%s", name, listResource)

	if err != nil {
		return err
	}

	for _, resource := range resources {
		resourceNode := node.AddNode(resource.Name)
		resourceNode.Config = resource.Value
	}

	return nil
}

// Add additional config to a node from another auth resource
// Useful for combining multiple resources
func AppendAuthState(resourceClient client.Client, name string, resource string, subResource string, node *config.Node) error {
	node = node.Children[resource]

	if node == nil {
		return fmt.Errorf("unable to append state. No child %s", name)
	}

	for childName, node := range node.Children {
		payload, err := resourceClient.Read("auth/%s/%s/%s/%s", name, resource, childName, subResource)

		if err != nil {
			return err
		}

		client.MergePayloads(node.Config, payload)
	}

	return nil
}

// Reads group states for an auth backend
func ReadAuthGroupsState(client client.Client, name string, node *config.Node) error {
	return ReadAuthResourceState(client, name, "groups", "groups", node)
}

// Reads user states for an auth backend
func ReadAuthUsersState(client client.Client, name string, node *config.Node) error {
	return ReadAuthResourceState(client, name, "users", "users", node)
}

// Reads role states for an auth backend
func ReadAuthRolesState(client client.Client, name string, node *config.Node) error {
	return ReadAuthResourceState(client, name, "roles", "roles", node)
}

// Reads role states for an auth backend, but with the singular name "role"
func ReadAuthRoleState(client client.Client, name string, node *config.Node) error {
	return ReadAuthResourceState(client, name, "role", "role", node)
}

func AppendAuthRoleIdState(client client.Client, name string, node *config.Node) error {
	return AppendAuthState(client, name, "role", "role-id", node)
}
