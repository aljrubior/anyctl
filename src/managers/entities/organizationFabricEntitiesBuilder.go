package entities

import "github.com/aljrubior/anyctl/clients/organizationRuntimeFabrics/response"

func NewOrganizationFabricEntitiesBuilder(response *[]response.OrganizationFabricResponse) *OrganizationFabricEntitiesBuilder {
	return &OrganizationFabricEntitiesBuilder{
		response,
	}
}

type OrganizationFabricEntitiesBuilder struct {
	response *[]response.OrganizationFabricResponse
}

func (this *OrganizationFabricEntitiesBuilder) Build() *[]OrganizationFabricEntity {

	var entities []OrganizationFabricEntity

	for _, v := range *this.response {
		entities = append(entities, OrganizationFabricEntity{
			v,
		})
	}

	return &entities
}
