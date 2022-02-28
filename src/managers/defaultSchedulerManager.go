package managers

import (
	"github.com/aljrubior/anyctl/managers/entities"
	"github.com/aljrubior/anyctl/managers/requests"
	"github.com/aljrubior/anyctl/services"
)

func NewDefaultSchedulerManager(deploymentManager DeploymentManager, schedulerService services.SchedulerService) DefaultSchedulerManager {
	return DefaultSchedulerManager{
		schedulerService:  schedulerService,
		deploymentManager: deploymentManager,
	}
}

type DefaultSchedulerManager struct {
	schedulerService  services.SchedulerService
	deploymentManager DeploymentManager
}

func (this DefaultSchedulerManager) GetSchedulers(ctx *entities.CurrentContextEntity, deploymentId string) (*[]entities.SchedulerEntity, error) {

	schedulers, err := this.schedulerService.GetSchedulers(ctx.OrganizationId, ctx.EnvironmentId, ctx.AuthorizationToken, deploymentId)

	if err != nil {
		return nil, err
	}

	return entities.NewSchedulerEntitiesBuilder(schedulers).Build(), nil
}

func (this DefaultSchedulerManager) FindSchedulerByFlowName(ctx *entities.CurrentContextEntity, deploymentId, flowName string) (*entities.SchedulerEntity, *[]entities.SchedulerEntity, error) {

	schedulers, err := this.GetSchedulers(ctx, deploymentId)

	if err != nil {
		return nil, nil, err
	}

	for _, v := range *schedulers {
		if v.FlowName == flowName {
			return &v, nil, nil
		}
	}

	return nil, schedulers, nil
}

func (this DefaultSchedulerManager) EnableScheduler(ctx *entities.CurrentContextEntity, deploymentId, flowName string, enabled bool) (*[]entities.SchedulerEntity, *[]entities.SchedulerEntity, error) {

	schedulers, err := this.GetSchedulers(ctx, deploymentId)

	if err != nil {
		return nil, nil, err
	}

	exists := false

	for _, v := range *schedulers {
		if v.FlowName == flowName {
			exists = true
			break
		}
	}

	if !exists {
		return nil, schedulers, nil
	}

	req := requests.NewSchedulerEnableRequestBuilder().AddScheduler(flowName, enabled).Build()

	resp, err := this.schedulerService.EnableSchedulers(ctx.OrganizationId, ctx.EnvironmentId, ctx.AuthorizationToken, deploymentId, req)

	if err != nil {
		return nil, nil, err
	}

	return entities.NewSchedulerEntitiesBuilder(resp).Build(), nil, nil
}

func (this DefaultSchedulerManager) EnableSchedulers(ctx *entities.CurrentContextEntity, deploymentId string, enabled bool) (*[]entities.SchedulerEntity, error) {

	schedulers, err := this.GetSchedulers(ctx, deploymentId)

	if err != nil {
		return nil, err
	}

	requestBuilder := requests.NewSchedulerEnableRequestBuilder()

	for _, v := range *schedulers {
		requestBuilder.AddScheduler(v.FlowName, enabled)
	}

	req := requestBuilder.Build()

	resp, err := this.schedulerService.EnableSchedulers(ctx.OrganizationId, ctx.EnvironmentId, ctx.AuthorizationToken, deploymentId, req)

	if err != nil {
		return nil, err
	}

	return entities.NewSchedulerEntitiesBuilder(resp).Build(), nil
}

func (this DefaultSchedulerManager) RunScheduler(ctx *entities.CurrentContextEntity, deploymentId, flowName string) (*[]entities.SchedulerEntity, error) {

	schedulers, err := this.GetSchedulers(ctx, deploymentId)

	if err != nil {
		return nil, err
	}

	exists := false

	for _, v := range *schedulers {
		if v.FlowName == flowName {
			exists = true
			break
		}
	}

	if !exists {
		return schedulers, nil
	}

	err = this.schedulerService.RunScheduler(ctx.OrganizationId, ctx.EnvironmentId, ctx.AuthorizationToken, deploymentId, flowName)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (this DefaultSchedulerManager) UnmanageScheduler(ctx *entities.CurrentContextEntity, deploymentId, flowName string) (*[]entities.SchedulerEntity, error) {

	schedulers, err := this.GetSchedulers(ctx, deploymentId)

	if err != nil {
		return nil, err
	}

	exists := false

	for _, v := range *schedulers {
		if v.FlowName == flowName {
			exists = true
			break
		}
	}

	if !exists {
		return schedulers, nil
	}

	err = this.schedulerService.DeleteScheduler(ctx.OrganizationId, ctx.EnvironmentId, ctx.AuthorizationToken, deploymentId, flowName)

	if err != nil {
		return nil, err
	}

	return nil, nil
}
