package entities

import (
	"github.com/aljrubior/anyctl/clients/privateSpaces/response"
)

func NewPrivateSpaceEntitiesBuilder(response *[]response.PrivateSpaceResponse) *PrivateSpaceEntitiesBuilder {
	return &PrivateSpaceEntitiesBuilder{
		response,
	}
}

type PrivateSpaceEntitiesBuilder struct {
	response *[]response.PrivateSpaceResponse
}

func (this *PrivateSpaceEntitiesBuilder) Build() *[]PrivateSpaceEntity {

	var entities []PrivateSpaceEntity

	for _, v := range *this.response {
		entities = append(entities, PrivateSpaceEntity{
			v,
		})
	}
	return &entities
}
