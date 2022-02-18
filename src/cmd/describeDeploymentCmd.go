package cmd

import (
	"github.com/aljrubior/anyctl/handlers"
	"github.com/spf13/cobra"
)

var describeDeploymentCmd = &cobra.Command{
	Use:     "describe",
	Aliases: []string{"d"},
	Short:   "Describe a deployment",
	Run: func(cmd *cobra.Command, args []string) {

		deploymentHandler := handlers.NewDeploymentHandler(DeploymentManager, ConfigManager, TargetManager)

		switch len(args) {
		case 1:
			if err := deploymentHandler.DescribeDeployment(args[0]); err != nil {
				Console.LogError(err)
			}
		default:
			println("Error: Unsupported option. Try with 'anyctl sharedspaces describe [fabric-name|fabric-id]'")
		}
	},
}

func init() {
	deploymentsCmd.AddCommand(describeDeploymentCmd)
}
