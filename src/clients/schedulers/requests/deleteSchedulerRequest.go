package requests

import (
	"fmt"
	"github.com/aljrubior/anyctl/clients"
	"github.com/aljrubior/anyctl/conf"
	"net/http"
)

func NewDeleteSchedulerRequest(
	config *conf.SchedulerClientConfig,
	bearerToken,
	organizationId,
	environmentId,
	deploymentId,
	flowName string) *DeleteSchedulerRequest {

	return &DeleteSchedulerRequest{
		config:         config,
		bearerToken:    bearerToken,
		organizationId: organizationId,
		environmentId:  environmentId,
		deploymentId:   deploymentId,
		flowName:       flowName,
	}
}

type DeleteSchedulerRequest struct {
	clients.BaseHttpRequest
	config         *conf.SchedulerClientConfig
	bearerToken    string
	organizationId string
	environmentId  string
	deploymentId   string
	flowName       string
}

func (this *DeleteSchedulerRequest) buildUri() string {

	protocol := this.config.Protocol
	host := this.config.Host
	path := fmt.Sprintf(this.config.SchedulerPath, this.organizationId, this.environmentId, this.deploymentId, this.flowName)

	return fmt.Sprintf("%s://%s/%s", protocol, host, path)
}

func (this *DeleteSchedulerRequest) Build() *http.Request {

	uri := this.buildUri()

	req, _ := http.NewRequest(http.MethodDelete, uri, nil)

	this.AddDefaultHeaders(req, this.organizationId, this.environmentId, this.bearerToken)

	return req
}
