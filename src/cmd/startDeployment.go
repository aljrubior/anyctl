package cmd

import (
	"github.com/aljrubior/anyctl/handlers"
	"github.com/aljrubior/anyctl/logger"
	"github.com/spf13/cobra"
)

var startDeploymentCmd = &cobra.Command{
	Use:   "start",
	Short: "Start deployment",
	Run: func(cmd *cobra.Command, args []string) {

		deploymentHandler := handlers.NewDeploymentHandler(DeploymentManager, ConfigManager, TargetManager)

		switch len(args) {
		case 1:
			if err := deploymentHandler.StartDeployment(args[0]); err != nil {
				logger.Error(err.Error())
			}
		default:
			Console.LogInvalidParameters()
		}
	},
}

func init() {
	deploymentsCmd.AddCommand(startDeploymentCmd)
}
