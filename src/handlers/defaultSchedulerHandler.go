package handlers

import (
	"fmt"
	"github.com/aljrubior/anyctl/errors"
	"github.com/aljrubior/anyctl/managers"
	"github.com/aljrubior/anyctl/managers/entities"
	"github.com/aljrubior/anyctl/utils"
)

func NewDefaultSchedulerHandler(
	configManager managers.ConfigManager,
	deploymentManager managers.DeploymentManager,
	schedulerManager managers.SchedulerManager) *DefaultSchedulerHandler {

	return &DefaultSchedulerHandler{
		configManager:     configManager,
		deploymentManager: deploymentManager,
		schedulerManager:  schedulerManager,
	}
}

type DefaultSchedulerHandler struct {
	configManager     managers.ConfigManager
	deploymentManager managers.DeploymentManager
	schedulerManager  managers.SchedulerManager
}

func (this DefaultSchedulerHandler) GetSchedulers(deploymentName string) error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	deployment, deployments, err := this.deploymentManager.FindDeploymentByName(ctx, deploymentName)

	if err != nil {
		return err
	}

	if deployment == nil {
		return this.ThrowDeploymentNotFoundError(deploymentName, deployments)
	}

	schedulers, err := this.schedulerManager.GetSchedulers(ctx, deployment.Id)

	if err != nil {
		return err
	}

	utils.PrintSchedulers(schedulers)

	return nil
}

func (this DefaultSchedulerHandler) GetScheduler(deploymentName, flowName string) error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	deployment, deployments, err := this.deploymentManager.FindDeploymentByName(ctx, deploymentName)

	if err != nil {
		return err
	}

	if deployment == nil {
		return this.ThrowDeploymentNotFoundError(deploymentName, deployments)
	}

	scheduler, schedulers, err := this.schedulerManager.FindSchedulerByFlowName(ctx, deployment.Id, flowName)

	if scheduler == nil {
		return this.ThrowSchedulerNotFoundError(flowName, schedulers)
	}

	utils.PrintScheduler(scheduler)

	return nil
}

func (this DefaultSchedulerHandler) EnableScheduler(deploymentName, flowName string, enabled bool) error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	deployment, deployments, err := this.deploymentManager.FindDeploymentByName(ctx, deploymentName)

	if err != nil {
		return err
	}

	if deployment == nil {
		return this.ThrowDeploymentNotFoundError(deploymentName, deployments)
	}

	scheduler, schedulers, err := this.schedulerManager.EnableScheduler(ctx, deployment.Id, flowName, enabled)

	if err != nil {
		return err
	}

	if scheduler == nil {
		return this.ThrowSchedulerNotFoundError(flowName, schedulers)
	}

	utils.PrintSchedulers(scheduler)

	return nil
}

func (this DefaultSchedulerHandler) RunScheduler(deploymentName, flowName string) error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	deployment, deployments, err := this.deploymentManager.FindDeploymentByName(ctx, deploymentName)

	if err != nil {
		return err
	}

	if deployment == nil {
		return this.ThrowDeploymentNotFoundError(deploymentName, deployments)
	}

	schedulers, err := this.schedulerManager.RunScheduler(ctx, deployment.Id, flowName)

	if err != nil {
		return err
	}

	if schedulers != nil {
		return this.ThrowSchedulerNotFoundError(flowName, schedulers)
	}

	return nil
}

func (this DefaultSchedulerHandler) UnmanageScheduler(deploymentName, flowName string) error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	deployment, deployments, err := this.deploymentManager.FindDeploymentByName(ctx, deploymentName)

	if err != nil {
		return err
	}

	if deployment == nil {
		return this.ThrowDeploymentNotFoundError(deploymentName, deployments)
	}

	schedulers, err := this.schedulerManager.UnmanageScheduler(ctx, deployment.Id, flowName)

	if err != nil {
		return err
	}

	if schedulers != nil {
		return this.ThrowSchedulerNotFoundError(flowName, schedulers)
	}

	println(fmt.Sprintf("Scheduler '%s' unmanaged.", flowName))
	return nil
}

func (this DefaultSchedulerHandler) ThrowDeploymentNotFoundError(deploymentName string, options *[]entities.DeploymentItemEntity) error {
	return errors.NewDeploymentNotFoundError(deploymentName, options)
}

func (this DefaultSchedulerHandler) ThrowSchedulerNotFoundError(flowName string, options *[]entities.SchedulerEntity) error {
	return errors.NewSchedulerNotFoundError(flowName, options)
}
