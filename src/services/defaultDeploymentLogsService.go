package services

import (
	"github.com/aljrubior/anyctl/clients/deploymentLogs"
	"github.com/aljrubior/anyctl/clients/deploymentLogs/response"
)

func NewDefaultDeploymentLogsService(logsClient deploymentLogs.DeploymentLogsClient) DefaultDeploymentLogsService {
	return DefaultDeploymentLogsService{
		logsClient: logsClient,
	}
}

type DefaultDeploymentLogsService struct {
	logsClient deploymentLogs.DeploymentLogsClient
}

func (this DefaultDeploymentLogsService) GetLogs(orgId, envId, token, deploymentId, specId string) (*[]response.DeploymentLogMessageResponse, error) {

	resp, err := this.logsClient.GetLogs(orgId, envId, token, deploymentId, specId)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
