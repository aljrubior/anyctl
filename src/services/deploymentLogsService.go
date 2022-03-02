package services

import "github.com/aljrubior/anyctl/clients/deploymentLogs/response"

type DeploymentLogsService interface {
	GetLogs(orgId, envId, token, deploymentId, specId string) (*[]response.DeploymentLogMessageResponse, error)
}
