package assets

import "github.com/aljrubior/anyctl/clients/assets/response"

type AssetClient interface {
	FindAssets(orgId, envId, token, assetName string) (*[]response.AssetResponse, error)
	FindLatestVersion(token, groupId, assetName string) (*response.AssetResponse, error)
	FindSpecificVersion(token, groupId, assetName, version string) (*response.AssetResponse, error)
	UploadArtifact(token, orgId, filePath, assetName, version string) (*response.AssetPublicationResponse, error)
}
