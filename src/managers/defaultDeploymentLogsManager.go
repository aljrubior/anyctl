package managers

import (
	"github.com/aljrubior/anyctl/managers/entities"
	"github.com/aljrubior/anyctl/services"
)

func NewDefaultDeploymentLogsManager(logsService services.DeploymentLogsService) DefaultDeploymentLogsManager {
	return DefaultDeploymentLogsManager{
		logsService: logsService,
	}
}

type DefaultDeploymentLogsManager struct {
	logsService services.DeploymentLogsService
}

func (this DefaultDeploymentLogsManager) GetLogs(ctx *entities.CurrentContextEntity, deploymentId, specId string) (*[]entities.DeploymentLogMessageEntity, error) {

	resp, err := this.logsService.GetLogs(ctx.OrganizationId, ctx.EnvironmentId, ctx.AuthorizationToken, deploymentId, specId)

	if err != nil {
		return nil, err
	}

	return entities.NewLogMessageEntitiesBuilder(resp).Build(), nil
}
