package cmd

import (
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Show or update the anyconfig file using subcommands like \"anyctl config set-environment environment-name\"",
	Run: func(cmd *cobra.Command, args []string) {

		println("Error: Unsupported option. Try with 'anyctl config --help'")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
