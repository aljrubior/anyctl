package entities

import (
	"github.com/aljrubior/anyctl/clients/assets/response"
)

func NewAssetUploadedEntityBuilder(response *response.AssetUploadedResponse) *AssetUploadedEntityBuilder {
	return &AssetUploadedEntityBuilder{
		response,
	}
}

type AssetUploadedEntityBuilder struct {
	response *response.AssetUploadedResponse
}

func (this AssetUploadedEntityBuilder) Build() *AssetUploadedEntity {
	return &AssetUploadedEntity{
		*this.response,
	}
}
