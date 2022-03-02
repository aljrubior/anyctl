package requests

import (
	"fmt"
	"github.com/aljrubior/anyctl/clients"
	"github.com/aljrubior/anyctl/conf"
	"net/http"
)

func NewGetDeploymentSpecsRequest(
	config *conf.DeploymentClientConfig,
	bearerToken,
	organizationId,
	environmentId,
	deploymentId string) *GetDeploymentSpecsRequest {

	return &GetDeploymentSpecsRequest{
		config:         config,
		bearerToken:    bearerToken,
		organizationId: organizationId,
		environmentId:  environmentId,
		deploymentId:   deploymentId,
	}
}

type GetDeploymentSpecsRequest struct {
	clients.BaseHttpRequest
	config         *conf.DeploymentClientConfig
	bearerToken    string
	organizationId string
	environmentId  string
	deploymentId   string
}

func (this *GetDeploymentSpecsRequest) buildUri() string {

	protocol := this.config.Protocol
	host := this.config.Host
	path := fmt.Sprintf(this.config.SpecsPath, this.organizationId, this.environmentId, this.deploymentId)

	return fmt.Sprintf("%s://%s/%s", protocol, host, path)
}

func (this *GetDeploymentSpecsRequest) Build() *http.Request {

	uri := this.buildUri()

	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	this.AddDefaultHeaders(req, this.organizationId, this.environmentId, this.bearerToken)

	return req
}
