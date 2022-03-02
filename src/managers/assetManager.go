package managers

import "github.com/aljrubior/anyctl/managers/entities"

type AssetManager interface {
	FindAssetByName(ctx *entities.CurrentContextEntity, assetName string) (*[]entities.AssetEntity, *[]entities.AssetEntity, error)
	FindLatestAssetByName(ctx *entities.CurrentContextEntity, assetName string) (*entities.AssetEntity, *[]entities.AssetEntity, error)
	FindAssetByRef(ctx *entities.CurrentContextEntity, assetRef string) (*entities.AssetEntity, error)
	FindLatestVersion(ctx *entities.CurrentContextEntity, groupId, assetName string) (*entities.AssetEntity, error)
	FindSpecificVersion(ctx *entities.CurrentContextEntity, groupId, assetName, version string) (*entities.AssetEntity, error)
	UploadAsset(ctx *entities.CurrentContextEntity, fromFilePath, withName, withVersion string) (*entities.AssetPublicationEntity, error)
}
