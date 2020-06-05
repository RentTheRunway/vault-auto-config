package cmd

import (
	pkg "github.com/RentTheRunway/vault-auto-config/internal/vault-auto-config"
	"github.com/spf13/cobra"
)

var (
	// flags
	inputDir string

	applyCmd = &cobra.Command{
		Use:   "apply",
		Short: "Applies the given vault configuration",
		Long:  "Applies the given vault configuration from the specified directory",
		RunE: func(cmd *cobra.Command, args []string) error {
			inputDir, err := cmd.Flags().GetString("input-dir")
			if err != nil {
				return err
			}

			vaultAutoConfig, err := pkg.NewVaultAutoConfig(vaultUrl, token)
			if err != nil {
				return err
			}

			return vaultAutoConfig.Apply(inputDir, secrets)
		},
	}
)

func init() {
	applyCmd.Flags().StringVarP(
		&inputDir,
		"input-dir",
		"i",
		"",
		"The input directory to apply vault configuration state from",
	)
	_ = applyCmd.MarkFlagRequired("input-dir")
	addVaultFlags(applyCmd)
	addSecretsFlag(applyCmd)
}
