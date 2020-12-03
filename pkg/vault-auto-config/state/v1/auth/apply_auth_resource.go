package auth

import (
	"fmt"

	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/client"
	"github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config/config"
)

// Applies config state for an auth backend
func ApplyAuthConfigState(node *config.Node, name string, client client.Client) error {
	node, ok := node.Children["config"]
	if !ok || node.Config == nil {
		return nil
	}

	return client.Write(node.Config, "auth/%s/config", name)
}

// Applies a generic auth resource state for an auth backend (e.g. groups, roles, etc.)
func ApplyAuthResourceState(node *config.Node, name string, resource string, client client.Client) error {
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
		} else if _, ok := node.Children[entry.Name]; !ok {
			remove = true
		}

		if remove {
			if err := client.Delete("auth/%s/%s/%s", name, resource, entry.Name); err != nil {
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

// Applied a field from a config to a sub resource
func ApplyAuthSubResourceState(node *config.Node, name string, resource string, configField string, subResource string, resourceClient client.Client) error {
	node = node.Children[resource]

	if node == nil {
		return fmt.Errorf("unable to apply subresource. No child %s", resource)
	}

	for childName, node := range node.Children {

		value, err := client.GetString(node.Config, configField)

		if err != nil {
			return err
		}

		m := map[string]interface{}{configField: value}

		if err := resourceClient.Write(m, "auth/%s/%s/%s/%s", name, resource, childName, subResource); err != nil {
			return err
		}
	}

	return nil
}

// Applies group states for an auth backend
func ApplyAuthGroupsState(node *config.Node, name string, client client.Client) error {
	return ApplyAuthResourceState(node, name, "groups", client)
}

// Applies user states for an auth backend
func ApplyAuthUsersState(node *config.Node, name string, client client.Client) error {
	return ApplyAuthResourceState(node, name, "users", client)
}

// Applies role states for an auth backend
func ApplyAuthRolesState(node *config.Node, name string, client client.Client) error {
	return ApplyAuthResourceState(node, name, "roles", client)
}

// Applies role states for an auth backend, but with the singular name "role"
func ApplyAuthRoleState(node *config.Node, name string, client client.Client) error {
	return ApplyAuthResourceState(node, name, "role", client)
}

func ApplyAuthApproleRoleIdState(node *config.Node, name string, client client.Client) error {
	return ApplyAuthSubResourceState(node, name, "role", "role_id", "role-id", client)
}

func ApplyAuthApproleSecretIdState(node *config.Node, name string, client client.Client) error {
	return ApplyAuthSubResourceState(node, name, "role", "secret_id", "custom-secret-id", client)
}
