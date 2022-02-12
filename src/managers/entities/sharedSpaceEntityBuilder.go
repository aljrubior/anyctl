package entities

import "github.com/aljrubior/anyctl/clients/sharedspaces/response"

func NewSharedSpaceEntityBuilder(response *response.SharedSpaceResponse) *SharedSpaceEntityBuilder {
	return &SharedSpaceEntityBuilder{
		response,
	}
}

type SharedSpaceEntityBuilder struct {
	response *response.SharedSpaceResponse
}

func (this *SharedSpaceEntityBuilder) Build() *SharedSpaceEntity {

	return &SharedSpaceEntity{
		*this.response,
	}
}
