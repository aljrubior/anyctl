package services

import (
	"encoding/json"
	"github.com/aljrubior/anyctl/clients/deployments"
	"github.com/aljrubior/anyctl/clients/deployments/response"
	"github.com/aljrubior/anyctl/managers/requests"
)

func NewDefaultDeploymentService(deploymentClient deployments.DeploymentClient) *DefaultDeploymentService {
	return &DefaultDeploymentService{
		deploymentClient: deploymentClient,
	}
}

type DefaultDeploymentService struct {
	deploymentClient deployments.DeploymentClient
}

func (this *DefaultDeploymentService) GetDeployments(orgId, envId, token string) (*[]response.DeploymentItem, error) {

	resp, err := this.deploymentClient.GetDeployments(orgId, envId, token)

	if err != nil {
		return nil, err
	}

	return &resp.Items, nil
}

func (this *DefaultDeploymentService) GetDeployment(orgId, envId, token, deploymentId string) (*response.DeploymentResponse, error) {

	resp, err := this.deploymentClient.GetDeployment(orgId, envId, token, deploymentId)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (this *DefaultDeploymentService) Deploy(orgId, envId, token string, request *requests.DeploymentRequest) (*response.DeploymentResponse, error) {

	body, err := json.Marshal(request)

	if err != nil {
		return nil, err
	}

	resp, err := this.deploymentClient.PostDeployment(orgId, envId, token, body)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (this *DefaultDeploymentService) UpdateDeployment(orgId, envId, token, deploymentId string, request *requests.DeploymentRequest) (*response.DeploymentResponse, error) {

	body, err := json.Marshal(request)

	if err != nil {
		return nil, err
	}

	resp, err := this.deploymentClient.PatchDeployment(orgId, envId, token, deploymentId, body)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (this *DefaultDeploymentService) DeleteDeployment(orgId, envId, token, deploymentId string) error {

	return this.deploymentClient.DeleteDeployment(orgId, envId, token, deploymentId)
}
