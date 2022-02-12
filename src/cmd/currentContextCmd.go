package cmd

import (
	"github.com/aljrubior/anyctl/handlers"
	"github.com/spf13/cobra"
)

var currentEnvironmentCmd = &cobra.Command{
	Use:   "current-environment",
	Short: "Displays the current environment.",
	Run: func(cmd *cobra.Command, args []string) {

		switch len(args) {
		case 0:
			configHandler := handlers.NewDefaultConfigHandler(*AccountManager, *ConfigManager)

			if err := configHandler.PrintCurrentContext(); err != nil {
				Console.LogError(err)
			}
		default:
			println("Error: Unsupported option. Try with 'anyctl config current-environment'")
		}
	},
}

func init() {
	configCmd.AddCommand(currentEnvironmentCmd)
}
