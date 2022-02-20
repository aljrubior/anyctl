package cmd

import (
	"github.com/aljrubior/anyctl/errors"
	"github.com/aljrubior/anyctl/handlers"
	"github.com/spf13/cobra"
)

var getHistoryCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"g"},
	Short:   "Retrieve the deployment history",
	Run: func(cmd *cobra.Command, args []string) {

		deploymentHistoryHandler := handlers.NewDefaultDeploymentHistoryHandler(DeploymentManager, ConfigManager)

		switch len(args) {
		case 1:
			if err := deploymentHistoryHandler.GetDeploymentHistory(args[0]); err != nil {
				errors.Catch(err).Println()
			}
		default:
			println("Error: Unsupported option. Try with 'anyctl runtimemanager deployments history get <deployment-name>'")
		}
	},
}

func init() {
	historyCmd.AddCommand(getHistoryCmd)
}
