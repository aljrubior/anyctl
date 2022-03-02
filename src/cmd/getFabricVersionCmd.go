package cmd

import (
	"github.com/aljrubior/anyctl/handlers"
	"github.com/aljrubior/anyctl/logger"
	"github.com/spf13/cobra"
)

var getFabricVersionsCmd = &cobra.Command{
	Use:     "get-versions",
	Aliases: []string{"get-version"},
	Short:   "List the version information for a Fabric",
	Run: func(cmd *cobra.Command, args []string) {

		fabricHandler := handlers.NewDefaultFabricHandler(ConfigManager, FabricManager, AccountManager)

		switch len(args) {
		case 1:
			if err := fabricHandler.GetVersions(args[0]); err != nil {
				logger.Error(err.Error())
			}
		default:
			println("Error: Unsupported option. Try with 'anyctl fabrics get-versions <fabric-name|fabric-id> --help'")
		}
	},
}

func init() {
	fabricsCmd.AddCommand(getFabricVersionsCmd)
}
