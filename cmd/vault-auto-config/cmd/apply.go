package cmd

import (
	pkg "github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config"
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
			vaultAutoConfig := pkg.NewVaultAutoConfig()
			return vaultAutoConfig.Apply(url, token, inputDir, secrets, additionalValues)
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
	addValuesFlag(applyCmd)
}
