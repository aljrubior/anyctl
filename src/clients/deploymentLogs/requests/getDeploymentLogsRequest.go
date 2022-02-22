package requests

import (
	"fmt"
	"github.com/aljrubior/anyctl/clients"
	"github.com/aljrubior/anyctl/conf"
	"net/http"
)

func NewGetLogsRequest(
	config *conf.DeploymentLogsClientConfig,
	bearerToken,
	organizationId,
	environmentId,
	deploymentId,
	specId string) *GetDeploymentLogsRequest {

	return &GetDeploymentLogsRequest{
		config:         config,
		bearerToken:    bearerToken,
		organizationId: organizationId,
		environmentId:  environmentId,
		deploymentId:   deploymentId,
		specId:         specId,
	}
}

type GetDeploymentLogsRequest struct {
	clients.BaseHttpRequest
	config         *conf.DeploymentLogsClientConfig
	bearerToken    string
	organizationId string
	environmentId  string
	deploymentId   string
	specId         string
}

func (this *GetDeploymentLogsRequest) buildUri() string {

	protocol := this.config.Protocol
	host := this.config.Host
	path := fmt.Sprintf(this.config.LogsPath, this.organizationId, this.environmentId, this.deploymentId, this.specId)
	return fmt.Sprintf("%s://%s/%s", protocol, host, path)
}

func (this *GetDeploymentLogsRequest) Build() *http.Request {

	uri := this.buildUri()

	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	this.AddDefaultHeaders(req, this.organizationId, this.environmentId, this.bearerToken)

	q := req.URL.Query()
	q.Add("descending", "true")
	req.URL.RawQuery = q.Encode()

	return req
}
