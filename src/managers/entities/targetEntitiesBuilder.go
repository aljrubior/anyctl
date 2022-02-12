package entities

import "github.com/aljrubior/anyctl/clients/targets/response"

func NewTargetEntitiesBuilder(response *[]response.TargetResponse) *TargetEntitiesBuilder {
	return &TargetEntitiesBuilder{
		response: response,
	}
}

type TargetEntitiesBuilder struct {
	response *[]response.TargetResponse
}

func (this TargetEntitiesBuilder) Build() *[]TargetEntity {
	var entities []TargetEntity

	for _, v := range *this.response {

		switch v.(type) {
		case *response.RuntimeFabricTargetResponse:

			if response, ok := v.(*response.RuntimeFabricTargetResponse); ok {
				entities = append(entities, &RuntimeFabricTargetEntity{*response})
				continue
			}

		case *response.StandaloneTargetResponse:

			if response, ok := v.(*response.StandaloneTargetResponse); ok {
				entities = append(entities, &StandaloneTargetEntity{*response})
				continue
			}
		}
	}

	return &entities
}
