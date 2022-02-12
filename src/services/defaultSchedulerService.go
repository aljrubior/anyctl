package services

import (
	"encoding/json"
	"github.com/aljrubior/anyctl/clients/schedulers"
	"github.com/aljrubior/anyctl/clients/schedulers/response"
	"github.com/aljrubior/anyctl/managers/requests"
)

func NewDefaultSchedulerService(schedulerClient schedulers.SchedulerClient) DefaultSchedulerService {
	return DefaultSchedulerService{
		schedulerClient: schedulerClient,
	}
}

type DefaultSchedulerService struct {
	schedulerClient schedulers.SchedulerClient
}

func (this DefaultSchedulerService) GetSchedulers(orgId, envId, token, deploymentId string) (*[]response.Scheduler, error) {

	resp, err := this.schedulerClient.GetSchedulers(orgId, envId, token, deploymentId)

	if err != nil {
		return nil, err
	}

	return &resp.Items, nil
}

func (this DefaultSchedulerService) EnableSchedulers(orgId, envId, token, deploymentId string, request *requests.SchedulerEnableRequest) (*[]response.Scheduler, error) {

	body, err := json.Marshal(request.Schedulers)

	if err != nil {
		return nil, err
	}

	resp, err := this.schedulerClient.PatchSchedulers(orgId, envId, token, deploymentId, body)

	if err != nil {
		return nil, err
	}

	return &resp.Items, nil
}

func (this DefaultSchedulerService) RunScheduler(orgId, envId, token, deploymentId, flowName string) error {

	return this.schedulerClient.PostScheduler(orgId, envId, token, deploymentId, flowName)
}

func (this DefaultSchedulerService) DeleteScheduler(orgId, envId, token, deploymentId, flowName string) error {

	return this.schedulerClient.DeleteScheduler(orgId, envId, token, deploymentId, flowName)
}
