package vault_auto_config

import (
	"errors"
	"github.com/RentTheRunway/vault-auto-config/internal/vault-auto-config/state"
	"os"
)

type VaultAutoConfig struct {
	url   string
	token string
}

// Creates a new vault configurator object
func NewVaultAutoConfig(url string, token string) (*VaultAutoConfig, error) {
	return &VaultAutoConfig{url: url, token: token}, nil
}

// Dumps the vault configuration state to a directory, overwriting if force is set to true
func (c *VaultAutoConfig) Dump(outputDir string, force bool) error {
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

	vault, err := state.NewVaultClient(c.url, c.token)
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
func (c *VaultAutoConfig) Apply(inputDir string, secrets string) error {
	vault, err := state.NewVaultClient(c.url, c.token)
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
