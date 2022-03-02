package entities

import "github.com/aljrubior/anyctl/clients/organizationRuntimeFabrics/response"

func NewOrganizationFabricEntityBuilder(response *response.OrganizationFabricResponse) *OrganizationFabricEntityBuilder {
	return &OrganizationFabricEntityBuilder{
		response,
	}
}

type OrganizationFabricEntityBuilder struct {
	response *response.OrganizationFabricResponse
}

func (this *OrganizationFabricEntityBuilder) Build() *OrganizationFabricEntity {

	return &OrganizationFabricEntity{
		*this.response,
	}
}
