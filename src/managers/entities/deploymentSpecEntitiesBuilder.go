package entities

import "github.com/aljrubior/anyctl/clients/deployments/response"

func NewDeploymentSpecEntitiesBuilder(response *[]response.DeploymentSpecResponse) *DeploymentSpecEntitiesBuilder {
	return &DeploymentSpecEntitiesBuilder{
		response,
	}
}

type DeploymentSpecEntitiesBuilder struct {
	response *[]response.DeploymentSpecResponse
}

func (this *DeploymentSpecEntitiesBuilder) Build() *[]DeploymentSpecEntity {

	var entities []DeploymentSpecEntity

	for _, v := range *this.response {
		entities = append(entities, DeploymentSpecEntity{
			v,
		})
	}
	return &entities
}
