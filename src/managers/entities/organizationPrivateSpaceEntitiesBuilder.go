package entities

import "github.com/aljrubior/anyctl/clients/organizationPrivateSpaces/response"

func NewOrganizationPrivateSpaceEntitiesBuilder(response *[]response.OrganizationPrivateSpaceResponse) *OrganizationPrivateSpaceEntitiesBuilder {
	return &OrganizationPrivateSpaceEntitiesBuilder{
		response,
	}
}

type OrganizationPrivateSpaceEntitiesBuilder struct {
	response *[]response.OrganizationPrivateSpaceResponse
}

func (this *OrganizationPrivateSpaceEntitiesBuilder) Build() *[]OrganizationPrivateSpaceEntity {

	var entities []OrganizationPrivateSpaceEntity

	for _, v := range *this.response {
		entities = append(entities, OrganizationPrivateSpaceEntity{
			v,
		})
	}
	return &entities
}
