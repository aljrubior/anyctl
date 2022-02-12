package entities

import (
	"github.com/aljrubior/anyctl/clients/privateSpaces/response"
)

func NewPrivateSpaceEntityBuilder(response *response.PrivateSpaceResponse) *PrivateSpaceEntityBuilder {
	return &PrivateSpaceEntityBuilder{
		response,
	}
}

type PrivateSpaceEntityBuilder struct {
	response *response.PrivateSpaceResponse
}

func (this *PrivateSpaceEntityBuilder) Build() *PrivateSpaceEntity {

	return &PrivateSpaceEntity{
		*this.response,
	}
}
