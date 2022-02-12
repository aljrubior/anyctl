package entities

import "github.com/aljrubior/anyctl/clients/fabrics/response"

func NewFabricEntitiesBuilder(response *[]response.FabricResponse) *FabricEntitiesBuilder {
	return &FabricEntitiesBuilder{
		response,
	}
}

type FabricEntitiesBuilder struct {
	response *[]response.FabricResponse
}

func (this *FabricEntitiesBuilder) Build() *[]FabricEntity {

	var entities []FabricEntity

	for _, v := range *this.response {
		entities = append(entities, FabricEntity{
			v,
		})
	}
	return &entities
}
