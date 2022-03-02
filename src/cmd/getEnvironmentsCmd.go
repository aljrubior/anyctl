package cmd

import (
	"github.com/aljrubior/anyctl/handlers"
	"github.com/spf13/cobra"
)

var getEnvironmentsCmd = &cobra.Command{
	Use:     "get-environments",
	Aliases: []string{"get-environment", "get-env"},
	Short:   "Displays the Anypoint environments available in the configuration file.",
	Run: func(cmd *cobra.Command, args []string) {

		switch len(args) {
		case 0:
			configHandler := handlers.NewDefaultConfigHandler(AccountManager, ConfigManager)

			if err := configHandler.GetEnvironments(); err != nil {
				Console.LogError(err)
			}
		default:
			println("Error: Unsupported option. Try with 'anyctl config get-environments'")
		}
	},
}

func init() {
	configCmd.AddCommand(getEnvironmentsCmd)
}
