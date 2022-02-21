package cmd

import (
	"github.com/aljrubior/anyctl/errors"
	"github.com/aljrubior/anyctl/handlers"
	"github.com/spf13/cobra"
)

var deploymentConfigurationFile *string

var applyDeploymentsCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply deployment configuration from a file",
	Run: func(cmd *cobra.Command, args []string) {

		applyDeploymentHandler := handlers.NewDefaultApplyDeploymentHandler(ConfigManager, DeployerManager, DeploymentManager, TargetManager, AssetManager)

		switch len(args) {
		case 0:
			if err := applyDeploymentHandler.Apply(*deploymentConfigurationFile); err != nil {
				errors.Catch(err).Println()
			}
		default:
			println("Error: Unsupported option. Try with 'anyctl runtimemanager deployments apply -f <deployment-file-path>'")
		}
	},
}

func init() {
	deploymentsCmd.AddCommand(applyDeploymentsCmd)

	deploymentConfigurationFile = applyDeploymentsCmd.Flags().StringP("file", "f", "", "File that contains the deployment configuration")
}
