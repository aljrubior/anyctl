package entities

import "github.com/aljrubior/anyctl/clients/organizationPrivateSpaces/response"

func NewOrganizationPrivateSpaceFabricEntitiesBuilder(response *[]response.OrganizationPrivateSpaceFabricResponse) *OrganizationPrivateSpaceFabricEntitiesBuilder {
	return &OrganizationPrivateSpaceFabricEntitiesBuilder{
		response,
	}
}

type OrganizationPrivateSpaceFabricEntitiesBuilder struct {
	response *[]response.OrganizationPrivateSpaceFabricResponse
}

func (this *OrganizationPrivateSpaceFabricEntitiesBuilder) Build() *[]OrganizationPrivateSpaceFabricEntity {

	var entities []OrganizationPrivateSpaceFabricEntity

	for _, v := range *this.response {
		entities = append(entities, OrganizationPrivateSpaceFabricEntity{
			v,
		})
	}
	return &entities
}
