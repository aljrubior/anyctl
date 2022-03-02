package deployments

import "github.com/aljrubior/anyctl/clients/deployments/response"

type DeploymentClient interface {
	GetDeployments(orgId, envId, token string) (*response.DeploymentsResponse, error)
	GetDeployment(orgId, envId, token, deploymentId string) (*response.DeploymentResponse, error)
	PostDeployment(orgId, envId, token string, body []byte) (*response.DeploymentResponse, error)
	PatchDeployment(orgId, envId, token, deploymentId string, body []byte) (*response.DeploymentResponse, error)
	DeleteDeployment(orgId, envId, token, deploymentId string) error

	GetDeploymentSpecs(orgId, envId, token, deploymentId string) (*[]response.DeploymentSpecResponse, error)
}
