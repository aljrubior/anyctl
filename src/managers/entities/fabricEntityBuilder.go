package entities

import "github.com/aljrubior/anyctl/clients/fabrics/response"

func NewFabricEntityBuilder(response *response.FabricResponse) *FabricEntityBuilder {
	return &FabricEntityBuilder{
		response,
	}
}

type FabricEntityBuilder struct {
	response *response.FabricResponse
}

func (this FabricEntityBuilder) Build() *FabricEntity {

	return &FabricEntity{
		*this.response,
	}
}
