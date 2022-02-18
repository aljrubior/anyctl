package cmd

import (
	"github.com/aljrubior/anyctl/errors"
	"github.com/aljrubior/anyctl/handlers"
	"github.com/spf13/cobra"
)

var getDeploymentsCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"g"},
	Short:   "Retrieve a list of deployments",
	Run: func(cmd *cobra.Command, args []string) {

		deploymentHandler := handlers.NewDeploymentHandler(DeploymentManager, ConfigManager, TargetManager)

		switch len(args) {
		case 0:
			if err := deploymentHandler.GetDeployments(); err != nil {
				errors.Catch(err).Println()
			}
		case 1:
			if err := deploymentHandler.FindDeploymentsContainsName(args[0]); err != nil {
				errors.Catch(err).Println()
			}
		default:
			println("Error: Unsupported option. Try with 'anyctl deployments get'")
		}
	},
}

func init() {
	deploymentsCmd.AddCommand(getDeploymentsCmd)
}
