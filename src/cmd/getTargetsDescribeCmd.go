package cmd

import (
	"github.com/aljrubior/anyctl/errors"
	"github.com/aljrubior/anyctl/handlers"
	"github.com/spf13/cobra"
)

var getTargetsDescribeCmd = &cobra.Command{
	Use:     "describe",
	Aliases: []string{"desc"},
	Short:   "Describe the target",
	Run: func(cmd *cobra.Command, args []string) {

		targetHandler := handlers.NewDefaultTargetHandler(ConfigManager, TargetManager)

		switch len(args) {
		case 1:
			if err := targetHandler.DescribeTarget(args[0]); err != nil {
				errors.Catch(err).Println()
			}
		default:
			println("Error: Unsupported option. Try with 'anyctl runtimemanager targets describe <target-name>'")
		}
	},
}

func init() {
	targetsCmd.AddCommand(getTargetsDescribeCmd)
}
