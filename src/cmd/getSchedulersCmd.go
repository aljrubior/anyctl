package cmd

import (
	"github.com/aljrubior/anyctl/errors"
	"github.com/aljrubior/anyctl/handlers"
	"github.com/spf13/cobra"
)

var getSchedulersCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"g"},
	Short:   "Retrieve a list of assets from a given name",
	Run: func(cmd *cobra.Command, args []string) {

		schedulerHandler := handlers.NewDefaultSchedulerHandler(ConfigManager, DeploymentManager, SchedulerManager)

		switch len(args) {
		case 1:
			if err := schedulerHandler.GetSchedulers(args[0]); err != nil {
				errors.Catch(err).Println()
			}
		case 2:
			if err := schedulerHandler.GetScheduler(args[0], args[1]); err != nil {
				errors.Catch(err).Println()
			}
		default:
			println("Error: Unsupported option. Try with 'anyctl runtimemanager scheduler get <deployment-name>'")
		}
	},
}

func init() {
	historyCmd.AddCommand(getSchedulersCmd)
}
