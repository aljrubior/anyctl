package entities

import "github.com/aljrubior/anyctl/clients/organizationRuntimeFabrics/response"

func NewFabricTargetEntitiesBuilder(response *[]response.FabricTargetResponse) *FabricTargetEntitiesBuilder {
	return &FabricTargetEntitiesBuilder{
		response,
	}
}

type FabricTargetEntitiesBuilder struct {
	response *[]response.FabricTargetResponse
}

func (this *FabricTargetEntitiesBuilder) Build() *[]FabricTargetEntity {

	var entities []FabricTargetEntity

	for _, v := range *this.response {
		entities = append(entities, FabricTargetEntity{
			v,
		})
	}

	return &entities
}
