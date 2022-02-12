package entities

import "github.com/aljrubior/anyctl/clients/sharedspaces/response"

func NewSharedSpaceEntitiesBuilder(response *[]response.SharedSpaceResponse) *SharedSpaceEntitiesBuilder {
	return &SharedSpaceEntitiesBuilder{
		response,
	}
}

type SharedSpaceEntitiesBuilder struct {
	response *[]response.SharedSpaceResponse
}

func (this *SharedSpaceEntitiesBuilder) Build() *[]SharedSpaceEntity {

	var entities []SharedSpaceEntity

	for _, v := range *this.response {
		entities = append(entities, SharedSpaceEntity{
			v,
		})
	}
	return &entities
}
