package cmd

import (
	"github.com/aljrubior/anyctl/errors"
	"github.com/aljrubior/anyctl/handlers"
	"github.com/spf13/cobra"
)

var runAssetRef *string
var targetName *string
var runtimeVersion *string

var runDeploymentCmd = &cobra.Command{
	Use:   "run",
	Short: "run",
	Run: func(cmd *cobra.Command, args []string) {

		deployerHandler := handlers.NewDefaultRunHandler(ConfigManager, DeployerManager)

		switch len(args) {
		case 1:
			if err := deployerHandler.Deploy(args[0], *runAssetRef, *targetName, *runtimeVersion); err != nil {
				errors.Catch(err).Println()
			}
		default:
			Console.LogInvalidParameters()
		}
	},
}

func init() {
	deploymentsCmd.AddCommand(runDeploymentCmd)

	runAssetRef = runDeploymentCmd.Flags().StringP("asset", "", "", "Asset")
	targetName = runDeploymentCmd.Flags().StringP("target-name", "", "", "Target")
	runtimeVersion = runDeploymentCmd.Flags().StringP("runtime-version", "", "", "Runtime")
}
