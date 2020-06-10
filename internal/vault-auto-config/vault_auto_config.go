package vault_auto_config

import (
	"errors"
	"github.com/RentTheRunway/vault-auto-config/internal/vault-auto-config/state"
	yaml2 "github.com/goccy/go-yaml"
	"os"
)

type VaultAutoConfig struct {
}

// Creates a new vault configurator object
func NewVaultAutoConfig() *VaultAutoConfig {
	return &VaultAutoConfig{}
}

// Dumps the vault configuration state to a directory, overwriting if force is set to true
func (c *VaultAutoConfig) Dump(url string, token string, outputDir string, force bool) error {
	if !force {
		isEmpty, err := IsEmptyDir(outputDir)
		if err != nil {
			return err
		}

		if !isEmpty {
			return errors.New("output directory is not empty")
		}
	} else {
		_ = os.RemoveAll(outputDir)
	}

	vault, err := state.NewVaultClient(url, token)
	if err != nil {
		return err
	}

	file, err := state.NewFileSystemClient(outputDir, "")
	if err != nil {
		return err
	}

	config, err := state.ReadState(vault)

	if err != nil {
		return err
	}

	return state.ApplyState(config, file)
}

// Applies vault configuration from a directory, optionally, with a secrets file to decrypt using sops
func (c *VaultAutoConfig) Apply(url string, token string, inputDir string, secrets string) error {
	vault, err := state.NewVaultClient(url, token)
	if err != nil {
		return err
	}

	file, err := state.NewFileSystemClient(inputDir, secrets)
	if err != nil {
		return err
	}

	config, err := state.ReadState(file)

	if err != nil {
		return err
	}

	return state.ApplyState(config, vault)
}

func (c *VaultAutoConfig) FileState(inputDir string, secrets string) (string, error) {
	client, err := state.NewFileSystemClient(inputDir, secrets)
	if err != nil {
		return "", err
	}

	configState, err := state.ReadState(client)
	if err != nil {
		return "", err
	}

	bytes, err := yaml2.Marshal(configState)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (c *VaultAutoConfig) VaultState(url string, token string) (string, error) {
	client, err := state.NewVaultClient(url, token)
	if err != nil {
		return "", err
	}

	configState, err := state.ReadState(client)
	if err != nil {
		return "", err
	}

	bytes, err := yaml2.Marshal(configState)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
