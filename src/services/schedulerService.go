package services

import (
	"github.com/aljrubior/anyctl/clients/schedulers/response"
	"github.com/aljrubior/anyctl/managers/requests"
)

type SchedulerService interface {
	GetSchedulers(orgId, envId, token, deploymentId string) (*[]response.Scheduler, error)
	EnableSchedulers(orgId, envId, token, deploymentId string, request *requests.SchedulerEnableRequest) (*[]response.Scheduler, error)
	RunScheduler(orgId, envId, token, deploymentId, flowName string) error
	DeleteScheduler(orgId, envId, token, deploymentId string, flowName string) error
}
