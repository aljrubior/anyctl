package entities

import "github.com/aljrubior/anyctl/clients/organizationPrivateSpaces/response"

func NewOrganizationPrivateSpaceEntityBuilder(response *response.OrganizationPrivateSpaceResponse) *OrganizationPrivateSpaceEntityBuilder {
	return &OrganizationPrivateSpaceEntityBuilder{
		response,
	}
}

type OrganizationPrivateSpaceEntityBuilder struct {
	response *response.OrganizationPrivateSpaceResponse
}

func (this *OrganizationPrivateSpaceEntityBuilder) Build() *OrganizationPrivateSpaceEntity {

	return &OrganizationPrivateSpaceEntity{
		*this.response,
	}
}
