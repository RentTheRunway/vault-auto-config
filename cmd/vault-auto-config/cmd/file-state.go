package cmd

import (
	"fmt"
	"github.com/RentTheRunway/vault-auto-config/internal/vault-auto-config/state"
	yaml2 "github.com/goccy/go-yaml"
	"github.com/spf13/cobra"
)

var (
	fileStateCmd = &cobra.Command{
		Use:   "file-state",
		Short: "Returns the current vault configuration state from the file system",
		Long:  "Returns the current vault configuration state from the file system",
		RunE: func(cmd *cobra.Command, args []string) error {
			inputDir, err := cmd.Flags().GetString("input-dir")
			if err != nil {
				return err
			}

			client, err := state.NewFileSystemClient(inputDir, secrets)
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
	fileStateCmd.Flags().StringVarP(
		&inputDir,
		"input-dir",
		"i",
		"",
		"The input directory to read vault configuration state from",
	)
	_ = fileStateCmd.MarkFlagRequired("input-dir")
	addSecretsFlag(fileStateCmd)
}
