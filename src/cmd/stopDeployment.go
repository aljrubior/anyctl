package cmd

import (
	"github.com/aljrubior/anyctl/handlers"
	"github.com/aljrubior/anyctl/logger"
	"github.com/spf13/cobra"
)

var stopDeploymentCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop deployment",
	Run: func(cmd *cobra.Command, args []string) {

		deploymentHandler := handlers.NewDeploymentHandler(*DeploymentManager, *ConfigManager, *TargetManager)

		switch len(args) {
		case 1:
			if err := deploymentHandler.StopDeployment(args[0]); err != nil {
				logger.Error(err.Error())
			}
		default:
			Console.LogInvalidParameters()
		}
	},
}

func init() {
	deploymentsCmd.AddCommand(stopDeploymentCmd)
}
