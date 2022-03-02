package services

import (
	"github.com/aljrubior/anyctl/clients/assets"
	"github.com/aljrubior/anyctl/clients/assets/response"
)

func NewDefaultAssetService(assetClient assets.AssetClient) DefaultAssetService {
	return DefaultAssetService{
		assetClient: assetClient,
	}
}

type DefaultAssetService struct {
	assetClient assets.AssetClient
}

func (this DefaultAssetService) FindAssets(orgId, envId, token, assetName string) (*[]response.AssetResponse, error) {
	return this.assetClient.FindAssets(orgId, envId, token, assetName)
}

func (this DefaultAssetService) FindLatestVersion(token, groupId, assetName string) (*response.AssetResponse, error) {
	return this.assetClient.FindLatestVersion(token, groupId, assetName)
}

func (this DefaultAssetService) FindSpecificVersion(token, groupId, assetName, version string) (*response.AssetResponse, error) {
	return this.assetClient.FindSpecificVersion(token, groupId, assetName, version)
}

func (this DefaultAssetService) UploadAsset(token, orgId, fromFilePath, withName, withVersion string) (*response.AssetPublicationResponse, error) {
	return this.assetClient.UploadArtifact(token, orgId, fromFilePath, withName, withVersion)
}
