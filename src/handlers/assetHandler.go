package handlers

type AssetHandler interface {
	FindAssets(assetName string) error
}
