package cmd

import (
	"github.com/aljrubior/anyctl/errors"
	"github.com/aljrubior/anyctl/handlers"
	"github.com/spf13/cobra"
)

var getPrivateSpaceFabricsCmd = &cobra.Command{
	Use:     "get-fabrics",
	Aliases: []string{"get-fabric"},
	Short:   "Lists all Fabrics for Private Space.",
	Run: func(cmd *cobra.Command, args []string) {

		privateSpaceHandler := handlers.NewDefaultPrivateSpaceHandler(*ConfigManager, *PrivateSpaceManager)

		switch len(args) {
		case 1:
			if err := privateSpaceHandler.GetFabrics(args[0]); err != nil {
				errors.Catch(err).Println()
			}
		default:
			println("Error: Unsupported option. Try with 'anyctl targets runtimefabrics nodes get <runtime-fabric-name>'")
		}
	},
}

func init() {
	privateSpacesCmd.AddCommand(getPrivateSpaceFabricsCmd)
}
