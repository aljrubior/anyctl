package managers

import (
	"github.com/aljrubior/anyctl/clients/assets/response"
	"github.com/aljrubior/anyctl/managers/entities"
	"github.com/aljrubior/anyctl/services"
	"github.com/aljrubior/anyctl/utils"
	"strings"
)

func NewDefaultAssetManager(assetService services.AssetService) *DefaultAssetManager {
	return &DefaultAssetManager{
		assetService: assetService,
	}
}

type DefaultAssetManager struct {
	assetService services.AssetService
}

func (this DefaultAssetManager) FindAssetByName(ctx *entities.CurrentContextEntity, assetName string) (*[]entities.AssetEntity, *[]entities.AssetEntity, error) {

	options, err := this.assetService.FindAssets(ctx.OrganizationId, ctx.EnvironmentId, ctx.AuthorizationToken, assetName)

	if err != nil {
		return nil, nil, err
	}

	var result []response.AssetResponse

	for _, v := range *options {
		if strings.Contains(v.Name, assetName) {
			result = append(result, v)
		}
	}
	
	return entities.NewAssetEntitiesBuilder(&result).Build(), entities.NewAssetEntitiesBuilder(options).Build(), nil
}

func (this DefaultAssetManager) FindLatestAssetByName(ctx *entities.CurrentContextEntity, assetName string) (*entities.AssetEntity, *[]entities.AssetEntity, error) {

	assets, options, err := this.FindAssetByName(ctx, assetName)

	if err != nil {
		return nil, nil, err
	}

	if assets == nil {
		return nil, options, nil
	}

	return &(*assets)[0], nil, nil
}

func (this DefaultAssetManager) FindAssetByNameAndVersion(ctx *entities.CurrentContextEntity, assetName, assetVersion string) (*entities.AssetEntity, *[]entities.AssetEntity, error) {

	assets, options, err := this.FindAssetByName(ctx, assetName)

	if err != nil {
		return nil, nil, err
	}

	if assets == nil {
		return nil, options, nil
	}

	for _, v := range *assets {
		println(v.Version)
		if v.Version == assetVersion {
			return &v, nil, nil
		}
	}

	return nil, assets, nil
}

func (this DefaultAssetManager) FindAssetByGroupAndNameAndVersion(ctx *entities.CurrentContextEntity, groupId, assetName, assetVersion string) (*entities.AssetEntity, *[]entities.AssetEntity, error) {

	assets, options, err := this.FindAssetByName(ctx, assetName)

	if err != nil {
		return nil, nil, err
	}

	if assets == nil {
		return nil, options, nil
	}

	for _, v := range *assets {
		if v.GroupId == groupId && v.Version == assetVersion {
			return &v, nil, nil
		}
	}

	return nil, options, nil
}

func (this DefaultAssetManager) FindAssetByRef(ctx *entities.CurrentContextEntity, assetRef string) (*entities.AssetEntity, error) {

	asset := strings.Split(assetRef, ":")

	switch len(asset) {
	case 1:
		return this.FindLatestVersion(ctx, ctx.OrganizationId, asset[0])
	case 2:
		if utils.IsValidUUID(asset[0]) {
			return this.FindLatestVersion(ctx, asset[0], asset[1])
		}

		return this.FindSpecificVersion(ctx, ctx.OrganizationId, asset[0], asset[1])
	case 3:
		return this.FindSpecificVersion(ctx, asset[0], asset[1], asset[2])
	default:
		return nil, nil
	}
}

func (this DefaultAssetManager) FindLatestVersion(ctx *entities.CurrentContextEntity, groupId, assetName string) (*entities.AssetEntity, error) {

	asset, err := this.assetService.FindLatestVersion(ctx.AuthorizationToken, groupId, assetName)

	if err != nil {
		return nil, err
	}

	return entities.NewAssetEntityBuilder(asset).Build(), nil
}

func (this DefaultAssetManager) FindSpecificVersion(ctx *entities.CurrentContextEntity, groupId, assetName, version string) (*entities.AssetEntity, error) {

	asset, err := this.assetService.FindSpecificVersion(ctx.AuthorizationToken, groupId, assetName, version)

	if err != nil {
		return nil, err
	}

	return entities.NewAssetEntityBuilder(asset).Build(), nil
}

func (this DefaultAssetManager) UploadAsset(ctx *entities.CurrentContextEntity, fromFilePath, withName, withVersion string) (*entities.AssetPublicationEntity, error) {

	response, err := this.assetService.UploadAsset(ctx.AuthorizationToken, ctx.OrganizationId, fromFilePath, withName, withVersion)

	if err != nil {
		return nil, err
	}

	return entities.NewAssetPublicationEntityBuilder(response).Build(), nil
}
