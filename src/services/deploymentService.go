package services

import (
	"github.com/aljrubior/anyctl/clients/deployments/response"
	"github.com/aljrubior/anyctl/managers/requests"
)

type DeploymentService interface {
	GetDeployments(orgId, envId, token string) (*[]response.DeploymentItem, error)
	GetDeployment(orgId, envId, token, deploymentId string) (*response.DeploymentResponse, error)
	Deploy(orgId, envId, token string, request *requests.DeploymentRequest) (*response.DeploymentResponse, error)
	UpdateDeployment(orgId, envId, token, deploymentId string, request *requests.DeploymentRequest) (*response.DeploymentResponse, error)
	DeleteDeployment(orgId, envId, token, deploymentId string) error
}
