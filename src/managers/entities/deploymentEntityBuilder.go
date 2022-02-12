package entities

import (
	"github.com/aljrubior/anyctl/clients/deployments/response"
)

func NewDeploymentEntityBuilder(response *response.DeploymentResponse) *DeploymentEntityBuilder {
	return &DeploymentEntityBuilder{
		response: response,
	}

}

type DeploymentEntityBuilder struct {
	response *response.DeploymentResponse
}

func (this DeploymentEntityBuilder) Build() *DeploymentEntity {

	return &DeploymentEntity{
		*this.response,
	}
}
