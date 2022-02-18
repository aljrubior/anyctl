package cmd

import (
	"github.com/aljrubior/anyctl/handlers"
	"github.com/aljrubior/anyctl/logger"
	"github.com/spf13/cobra"
)

var getFabricsCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"g"},
	Short:   "Retrieve a list of fabrics",
	Run: func(cmd *cobra.Command, args []string) {

		fabricHandler := handlers.NewDefaultFabricHandler(ConfigManager, FabricManager, AccountManager)

		switch len(args) {
		case 1:
			if err := fabricHandler.GetFabrics(args[0]); err != nil {
				logger.Error(err.Error())
			}
		default:
			println("Error: Unsupported option. Try with 'anyctl fabrics get <fabric-name> --help'")
		}
	},
}

func init() {
	fabricsCmd.AddCommand(getFabricsCmd)
}
