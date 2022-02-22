package entities

import "github.com/aljrubior/anyctl/clients/deploymentLogs/response"

func NewLogMessageEntitiesBuilder(response *[]response.DeploymentLogMessageResponse) *DeploymentLogMessageEntitiesBuilder {
	return &DeploymentLogMessageEntitiesBuilder{
		response: response,
	}
}

type DeploymentLogMessageEntitiesBuilder struct {
	response *[]response.DeploymentLogMessageResponse
}

func (this *DeploymentLogMessageEntitiesBuilder) Build() *[]DeploymentLogMessageEntity {

	var entities []DeploymentLogMessageEntity

	for _, v := range *this.response {
		entities = append(entities, DeploymentLogMessageEntity{
			v,
		})
	}

	return &entities
}
