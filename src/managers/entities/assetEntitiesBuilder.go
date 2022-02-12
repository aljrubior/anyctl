package entities

import "github.com/aljrubior/anyctl/clients/assets/response"

func NewAssetEntitiesBuilder(assets *[]response.AssetResponse) *AssetEntitiesBuilder {
	return &AssetEntitiesBuilder{
		assets: assets,
	}
}

type AssetEntitiesBuilder struct {
	assets *[]response.AssetResponse
}

func (this AssetEntitiesBuilder) Build() *[]AssetEntity {

	var result []AssetEntity

	for _, v := range *this.assets {
		result = append(result, AssetEntity{
			v,
		})
	}
	return &result
}
