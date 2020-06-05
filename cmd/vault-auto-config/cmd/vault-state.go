package cmd

import (
	"fmt"
	"github.com/RentTheRunway/vault-auto-config/internal/vault-auto-config/state"
	yaml2 "github.com/goccy/go-yaml"
	"github.com/spf13/cobra"
)

var (
	vaultStateCmd = &cobra.Command{
		Use:   "vault-state",
		Short: "Returns the current vault configuration state from vault",
		Long:  "Returns the current vault configuration state from vault",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := state.NewVaultClient(vaultUrl, token)
			if err != nil {
				return err
			}

			configState, err := state.ReadState(client)
			if err != nil {
				return err
			}

			bytes, err := yaml2.Marshal(configState)
			if err != nil {
				return err
			}

			fmt.Println(string(bytes))

			return nil
		},
	}
)

func init() {
	addVaultFlags(vaultStateCmd)
}
