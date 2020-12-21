package cmd

import (
	"fmt"
	pkg "github.com/RentTheRunway/vault-auto-config/pkg/vault-auto-config"
	"github.com/spf13/cobra"
)

var (
	fileStateCmd = &cobra.Command{
		Use:   "file-state",
		Short: "Returns the current vault configuration state from the file system",
		Long:  "Returns the current vault configuration state from the file system",
		RunE: func(cmd *cobra.Command, args []string) error {
			vaultAutoConfig := pkg.NewVaultAutoConfig()
			fmt.Println(vaultAutoConfig.FileState(inputDir, secrets, values))
			return nil
		},
	}
)

func init() {
	fileStateCmd.Flags().StringVarP(
		&inputDir,
		"input-dir",
		"i",
		"",
		"The input directory to read vault configuration state from",
	)
	_ = fileStateCmd.MarkFlagRequired("input-dir")
	addSecretsFlag(fileStateCmd)
	addValuesFlag(fileStateCmd)
}
