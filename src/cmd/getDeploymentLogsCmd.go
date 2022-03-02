package cmd

import (
	"github.com/aljrubior/anyctl/errors"
	"github.com/aljrubior/anyctl/handlers"
	"github.com/spf13/cobra"
)

var getDeploymentLogsCmd = &cobra.Command{
	Use:     "logs",
	Aliases: []string{"log"},
	Short:   "Retrieve logs of a deployment",
	Run: func(cmd *cobra.Command, args []string) {

		deploymentLogsHandler := handlers.NewDefaultDeploymentLogsHandler(DeploymentManager, ConfigManager, DeploymentLogsManager)

		switch len(args) {
		case 1:
			if err := deploymentLogsHandler.GetLogs(args[0]); err != nil {
				errors.Catch(err).Println()
			}
		default:
			println("Error: Unsupported option. Try with 'anyctl runtimemanager deployments logs <deployment-name>'")
		}
	},
}

func init() {
	deploymentsCmd.AddCommand(getDeploymentLogsCmd)
}
