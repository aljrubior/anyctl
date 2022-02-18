package cmd

import (
	"github.com/aljrubior/anyctl/errors"
	"github.com/aljrubior/anyctl/handlers"
	"github.com/spf13/cobra"
)

var assetFile *string
var assetName *string
var assetVersion *string

var uploadAssetCmd = &cobra.Command{
	Use:   "upload",
	Short: "upload",
	Run: func(cmd *cobra.Command, args []string) {

		assetHandler := handlers.NewDefaultAssetHandler(AssetManager, ConfigManager)

		switch len(args) {
		case 0:
			if err := assetHandler.UploadAsset(*assetFile, *assetName, *assetVersion); err != nil {
				errors.Catch(err).Println()
			}
		default:
			Console.LogInvalidParameters()
		}
	},
}

func init() {
	assetsCmd.AddCommand(uploadAssetCmd)

	assetFile = uploadAssetCmd.Flags().StringP("file", "", "", "File path")
	assetName = uploadAssetCmd.Flags().StringP("name", "", "", "Asset name")
	assetVersion = uploadAssetCmd.Flags().StringP("version", "", "", "Asset version")
}
