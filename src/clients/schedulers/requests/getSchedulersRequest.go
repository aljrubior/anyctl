package requests

import (
	"fmt"
	"github.com/aljrubior/anyctl/clients"
	"github.com/aljrubior/anyctl/conf"
	"net/http"
)

func NewGetSchedulersRequest(
	config *conf.SchedulerClientConfig,
	bearerToken,
	organizationId,
	environmentId,
	deploymentId string) *GetSchedulersRequest {

	return &GetSchedulersRequest{
		config:         config,
		bearerToken:    bearerToken,
		organizationId: organizationId,
		environmentId:  environmentId,
		deploymentId:   deploymentId,
	}
}

type GetSchedulersRequest struct {
	clients.BaseHttpRequest
	config         *conf.SchedulerClientConfig
	bearerToken    string
	organizationId string
	environmentId  string
	deploymentId   string
}

func (this *GetSchedulersRequest) buildUri() string {

	protocol := this.config.Protocol
	host := this.config.Host
	path := fmt.Sprintf(this.config.SchedulersPath, this.organizationId, this.environmentId, this.deploymentId)

	return fmt.Sprintf("%s://%s/%s", protocol, host, path)
}

func (this *GetSchedulersRequest) Build() *http.Request {

	uri := this.buildUri()

	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	this.AddDefaultHeaders(req, this.organizationId, this.environmentId, this.bearerToken)

	return req
}
