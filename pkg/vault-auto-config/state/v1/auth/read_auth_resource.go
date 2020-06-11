package auth

import (
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
func ReadAuthResourceState(client client.Client, name string, listResource string, resource string, node *config.Node) error {
	node = node.AddNode(resource)

	resources, err := client.List("auth/%s/%s", name, listResource)

	if err != nil {
		return err
	}

	for _, resource := range resources {
		resourceNode := node.AddNode(resource.Name)
		resourceNode.Config = resource.Value
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
