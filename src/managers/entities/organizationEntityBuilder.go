package entities

import "github.com/aljrubior/anyctl/clients/accounts/response"

func NewOrganizationEntityBuilder(organization *response.OrganizationResponse) *OrganizationEntityBuilder {
	return &OrganizationEntityBuilder{
		*organization,
	}
}

type OrganizationEntityBuilder struct {
	organization response.OrganizationResponse
}

func (this OrganizationEntityBuilder) Build() *OrganizationEntity {

	return &OrganizationEntity{
		this.organization,
	}
}
