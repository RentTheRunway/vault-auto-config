package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// common flags
	vaultUrl string
	token    string
	verbose  bool
	secrets string

	rootCmd = &cobra.Command{
		Use:   "vault-auto-config",
		Short: "Vault automatic configurator",
		Long: `Vault automatic configurator allows you to manage your vault configuration as code by structuring resources
in a directory structure that mimics the vault api`,
		SilenceUsage: true,
	}
)

func Execute() {
	_ = rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(
		&verbose,
		"verbose",
		"v",
		false,
		"Enable verbose logging",
	)
	rootCmd.AddCommand(dumpCmd)
	rootCmd.AddCommand(applyCmd)
	rootCmd.AddCommand(vaultStateCmd)
	rootCmd.AddCommand(fileStateCmd)
}

func addVaultFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(
		&vaultUrl,
		"url",
		"u",
		"http://127.0.0.1:8200",
		"The url of the vault api server",
	)
	cmd.Flags().StringVarP(
		&token,
		"token",
		"t",
		"",
		"The vault token to authenticate with",
	)
	_ = cmd.MarkFlagRequired("url")
	_ = cmd.MarkFlagRequired("token")
}
func addSecretsFlag(cmd *cobra.Command) {
	cmd.Flags().StringVarP(
		&secrets,
		"secrets",
		"s",
		"",
		"A secrets yaml file encrypted with sops to pass in for go template values",
	)
}