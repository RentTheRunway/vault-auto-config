package cmd

import (
	"fmt"
	pkg "github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config"
	"github.com/spf13/cobra"
)

var (
	vaultStateCmd = &cobra.Command{
		Use:   "vault-state",
		Short: "Returns the current vault configuration state from vault",
		Long:  "Returns the current vault configuration state from vault",
		RunE: func(cmd *cobra.Command, args []string) error {
			vaultAutoConfig := pkg.NewVaultAutoConfig()
			fmt.Println(vaultAutoConfig.VaultState(url, token))
			return nil
		},
	}
)

func init() {
	addVaultFlags(vaultStateCmd)
}
