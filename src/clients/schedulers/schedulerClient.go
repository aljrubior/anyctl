package schedulers

import "github.com/aljrubior/anyctl/clients/schedulers/response"

type SchedulerClient interface {
	GetSchedulers(orgId, envId, token, deploymentId string) (*response.SchedulersResponse, error)
	PatchSchedulers(orgId, envId, token, deploymentId string, body []byte) (*response.SchedulersResponse, error)
	PostScheduler(orgId, envId, token, deploymentId string, flowName string) error
	DeleteScheduler(orgId, envId, token, deploymentId string, flowName string) error
}
