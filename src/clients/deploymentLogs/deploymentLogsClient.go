package deploymentLogs

import "github.com/aljrubior/anyctl/clients/deploymentLogs/response"

type DeploymentLogsClient interface {
	GetLogs(orgId, envId, token, deploymentId, specId string) (*[]response.DeploymentLogMessageResponse, error)
}
