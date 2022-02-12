package response

type AssetUploadedResponse struct {
	OrganizationId string `json:"organizationId"`
	GroupId        string `json:"groupId"`
	AssetId        string `json:"assetId"`
	Version        string `json:"version"`
	Name           string `json:"name"`
	Type           string `json:"type"`
}
