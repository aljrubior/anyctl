package services

import "github.com/aljrubior/anyctl/clients/assets/response"

type AssetService interface {
	FindAssets(orgId, envId, token, assetName string) (*[]response.AssetResponse, error)
	FindLatestVersion(token, groupId, assetName string) (*response.AssetResponse, error)
	FindSpecificVersion(token, groupId, assetName, version string) (*response.AssetResponse, error)
	UploadAsset(token, orgId, fromFilePath, withName, withVersion string) (*response.AssetPublicationResponse, error)
}
