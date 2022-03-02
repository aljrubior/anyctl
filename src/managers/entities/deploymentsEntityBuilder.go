package entities

import "github.com/aljrubior/anyctl/clients/deployments/response"

func NewDeploymentsEntityBuilder(response *[]response.DeploymentItem) *DeploymentsEntityBuilder {

	return &DeploymentsEntityBuilder{
		response: response,
	}
}

type DeploymentsEntityBuilder struct {
	response *[]response.DeploymentItem
}

func (this DeploymentsEntityBuilder) Build() *[]DeploymentItemEntity {

	var deployments []DeploymentItemEntity

	for _, v := range *this.response {
		deployments = append(deployments, DeploymentItemEntity{
			v,
		})
	}

	return &deployments
}
