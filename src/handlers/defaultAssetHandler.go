package handlers

import (
	"fmt"
	errors2 "github.com/aljrubior/anyctl/errors"
	"github.com/aljrubior/anyctl/managers"
	"github.com/aljrubior/anyctl/managers/entities"
	"github.com/aljrubior/anyctl/printers"
)

func NewDefaultAssetHandler(assetManager managers.AssetManager, configManager managers.ConfigManager) *DefaultAssetHandler {
	return &DefaultAssetHandler{
		assetManager:  assetManager,
		configManager: configManager,
	}

}

type DefaultAssetHandler struct {
	AssetHandler

	assetManager  managers.AssetManager
	configManager managers.ConfigManager
}

func (this DefaultAssetHandler) FindAssets(assetName string) error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	assets, options, err := this.assetManager.FindAssetByName(ctx, assetName)

	if err != nil {
		return err
	}

	if assets == nil {
		return this.ThrowAssetNotFoundError(assetName, options)
	}

	printers.NewAssetsPrinter(assets).Print()

	return nil
}

func (this DefaultAssetHandler) UploadAsset(fromFilePath, withName, withVersion string) error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	asset, err := this.assetManager.UploadAsset(ctx, fromFilePath, withName, withVersion)

	if err != nil {
		return err
	}

	println(fmt.Sprintf("Asset '%s' v%s created.", withName, withVersion))
	println(asset.PublicationStatusLink)

	return nil
}

func (this DefaultAssetHandler) ThrowAssetNotFoundError(assetName string, assets *[]entities.AssetEntity) error {
	return errors2.NewAssetNotFoundError(assetName, assets)
}
