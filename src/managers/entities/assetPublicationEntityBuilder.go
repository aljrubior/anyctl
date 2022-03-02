package entities

import "github.com/aljrubior/anyctl/clients/assets/response"

func NewAssetPublicationEntityBuilder(response *response.AssetPublicationResponse) *AssetPublicationEntityBuilder {
	return &AssetPublicationEntityBuilder{
		response,
	}
}

type AssetPublicationEntityBuilder struct {
	response *response.AssetPublicationResponse
}

func (this AssetPublicationEntityBuilder) Build() *AssetPublicationEntity {
	return &AssetPublicationEntity{
		*this.response,
	}
}
