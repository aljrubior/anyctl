package cmd

import (
	"github.com/aljrubior/anyctl/errors"
	"github.com/aljrubior/anyctl/handlers"
	"github.com/spf13/cobra"
)

var replicas *int

var scaleDeploymentCmd = &cobra.Command{
	Use:   "scale",
	Short: "Scale deployment replicas",
	Run: func(cmd *cobra.Command, args []string) {

		deploymentHandler := handlers.NewDeploymentHandler(*DeploymentManager, *ConfigManager, *TargetManager)

		switch len(args) {
		case 1:
			if err := deploymentHandler.ScaleDeployment(args[0], *replicas); err != nil {
				errors.Catch(err).Println()
			}
		default:
			Console.LogInvalidParameters()
		}
	},
}

func init() {
	deploymentsCmd.AddCommand(scaleDeploymentCmd)
	replicas = scaleDeploymentCmd.Flags().IntP("replicas", "", -1, "Desired number of replicas")
}
