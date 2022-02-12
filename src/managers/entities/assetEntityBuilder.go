package entities

import "github.com/aljrubior/anyctl/clients/assets/response"

func NewAssetEntityBuilder(response *response.AssetResponse) *AssetEntityBuilder {
	return &AssetEntityBuilder{
		response,
	}
}

type AssetEntityBuilder struct {
	response *response.AssetResponse
}

func (this *AssetEntityBuilder) Build() *AssetEntity {

	return &AssetEntity{
		*this.response,
	}
}
