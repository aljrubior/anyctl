package managers

import "github.com/aljrubior/anyctl/managers/entities"

type SchedulerManager interface {
	GetSchedulers(ctx *entities.CurrentContextEntity, deploymentId string) (*[]entities.SchedulerEntity, error)
	FindSchedulerByFlowName(ctx *entities.CurrentContextEntity, deploymentId, flowName string) (*entities.SchedulerEntity, *[]entities.SchedulerEntity, error)
	EnableScheduler(ctx *entities.CurrentContextEntity, deploymentId, flowName string, enabled bool) (*[]entities.SchedulerEntity, *[]entities.SchedulerEntity, error)
	EnableSchedulers(ctx *entities.CurrentContextEntity, deploymentId string, enabled bool) (*[]entities.SchedulerEntity, error)
	RunScheduler(ctx *entities.CurrentContextEntity, deploymentId, flowName string) (*[]entities.SchedulerEntity, error)
	UnmanageScheduler(ctx *entities.CurrentContextEntity, deploymentId, flowName string) (*[]entities.SchedulerEntity, error)
}
