package cmd

import (
	"github.com/aljrubior/anyctl/errors"
	"github.com/aljrubior/anyctl/handlers"
	"github.com/spf13/cobra"
)

var getAssetsCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"g"},
	Short:   "Retrieve a list of assets from a given name",
	Run: func(cmd *cobra.Command, args []string) {

		assetHandler := handlers.NewDefaultAssetHandler(*AssetManager, *ConfigManager)

		switch len(args) {
		case 1:
			if err := assetHandler.FindAssets(args[0]); err != nil {
				errors.Catch(err).Println()
			}
		default:
			println("Error: Unsupported option. Try with 'anyctl assets get <asset-name>'")
		}
	},
}

func init() {
	assetsCmd.AddCommand(getAssetsCmd)
}
