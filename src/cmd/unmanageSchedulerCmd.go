package cmd

import (
	"github.com/aljrubior/anyctl/errors"
	"github.com/aljrubior/anyctl/handlers"
	"github.com/spf13/cobra"
)

var unmanageSchedulerCmd = &cobra.Command{
	Use:   "unmanage",
	Short: "Enable a scheduler in a specified deployment",
	Run: func(cmd *cobra.Command, args []string) {

		schedulerHandler := handlers.NewDefaultSchedulerHandler(*ConfigManager, *DeploymentManager, *SchedulerManager)

		// println(args[0], args[1], args[2], *setSchedulerEnabled)
		switch len(args) {
		case 2:

			if err := schedulerHandler.UnmanageScheduler(args[0], args[1]); err != nil {
				errors.Catch(err).Println()
				return
			}
		default:
			Console.LogInvalidParameters()
		}
	},
}

func init() {
	schedulersCmd.AddCommand(unmanageSchedulerCmd)
}
