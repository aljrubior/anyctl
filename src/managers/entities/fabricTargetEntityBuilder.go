package entities

import "github.com/aljrubior/anyctl/clients/organizationRuntimeFabrics/response"

func NewFabricTargetEntityBuilder(response *response.FabricTargetResponse) *FabricTargetEntityBuilder {
	return &FabricTargetEntityBuilder{
		response,
	}
}

type FabricTargetEntityBuilder struct {
	response *response.FabricTargetResponse
}

func (this *FabricTargetEntityBuilder) Build() *FabricTargetEntity {

	return &FabricTargetEntity{
		*this.response,
	}
}
