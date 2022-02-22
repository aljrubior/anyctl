package managers

import "github.com/aljrubior/anyctl/managers/entities"

type DeploymentLogsManager interface {
	GetLogs(ctx *entities.CurrentContextEntity, deploymentId, specId string) (*[]entities.DeploymentLogMessageEntity, error)
}
