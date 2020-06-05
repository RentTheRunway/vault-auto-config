package cmd

import (
	pkg "github.com/RentTheRunway/vault-auto-config/internal/vault-auto-config"
	"github.com/spf13/cobra"
)

var (
	// flags
	outputDir string
	force     bool

	dumpCmd = &cobra.Command{
		Use:   "dump",
		Short: "Dumps the current vault configuration",
		Long:  "Dumps the current vault configuration into the specified directory",
		RunE: func(cmd *cobra.Command, args []string) error {
			outputDir, err := cmd.Flags().GetString("output-dir")
			if err != nil {
				return err
			}

			vaultAutoConfig, err := pkg.NewVaultAutoConfig(vaultUrl, token)
			if err != nil {
				return err
			}

			return vaultAutoConfig.Dump(outputDir, force)
		},
	}
)

func init() {
	dumpCmd.Flags().StringVarP(
		&outputDir,
		"output-dir",
		"o",
		"",
		"The output directory to dump vault configuration state to",
	)
	_ = dumpCmd.MarkFlagRequired("output-dir")
	dumpCmd.Flags().BoolVarP(
		&force,
		"force",
		"f",
		false,
		"Forces dumping of state, overwriting the output directory",
	)

	addVaultFlags(dumpCmd)
}
