package handlers

import "github.com/aljrubior/anyctl/managers/entities"

type SchedulerHandler interface {
	GetSchedulers(deploymentName string) error
	GetScheduler(deploymentName, flowName string) error
	EnableScheduler(deploymentName, flowName string, enabled bool) error
	RunScheduler(deploymentName, flowName string) error
	UnmanageScheduler(deploymentName, flowName string) error

	ThrowDeploymentNotFoundError(deploymentName string, options *[]entities.DeploymentItemEntity) error
	ThrowSchedulerNotFoundError(flowName string, options *[]entities.SchedulerEntity) error
}
