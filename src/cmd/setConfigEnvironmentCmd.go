package cmd

import (
	"github.com/aljrubior/anyctl/errors"
	"github.com/aljrubior/anyctl/handlers"
	"github.com/spf13/cobra"
)

var setConfigEnvironmentCmd = &cobra.Command{
	Use:   "set-environment",
	Short: "Make specified environment active",
	Long:  `This command makes active the environment specified in <environment-name>`,
	Run: func(cmd *cobra.Command, args []string) {

		configHandler := handlers.NewDefaultConfigHandler(AccountManager, ConfigManager)

		if len(args) == 1 {
			environmentName := args[0]

			if err := configHandler.SetCurrentEnvironment(environmentName); err != nil {
				errors.Catch(err).Println()
				return
			}

			anyconfigPath, err := ConfigManager.GetAnyConfigFilePath()

			if err != nil {
				Console.LogError(err)
				return
			}

			Console.LogUpdatedCurrentEnvironmentSuccess(environmentName, anyconfigPath)
		} else {
			println("Error: Unsupported option. Try with 'anyctl config set-environment <environment-name>'")
		}
	},
}

func init() {
	configCmd.AddCommand(setConfigEnvironmentCmd)
}
