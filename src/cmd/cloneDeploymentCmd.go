package cmd

import (
	"github.com/aljrubior/anyctl/errors"
	"github.com/aljrubior/anyctl/handlers"
	"github.com/spf13/cobra"
)

var cloneWithName *string
var cloneToTargetName *string
var cloneToEnvironment *string

var cloneDeploymentCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone a deployment",
	Run: func(cmd *cobra.Command, args []string) {

		migrationHandler := handlers.NewDefaultDeploymentMigrationHandler(DeployerManager, ConfigManager, DeploymentManager, TargetManager, AccountManager)

		switch len(args) {
		case 1:
			if err := migrationHandler.Clone(args[0], *cloneWithName, *cloneToTargetName, *cloneToEnvironment); err != nil {
				errors.Catch(err).Println()
			}
		default:
			Console.LogInvalidParameters()
		}
	},
}

func init() {
	deploymentsCmd.AddCommand(cloneDeploymentCmd)
	cloneWithName = cloneDeploymentCmd.Flags().StringP("with-name", "", "", "Deployment name")
	cloneToTargetName = cloneDeploymentCmd.Flags().StringP("to-target-name", "", "", "Target destination ")
	cloneToEnvironment = cloneDeploymentCmd.Flags().StringP("to-environment-name", "", "", "Environment destination")
}
