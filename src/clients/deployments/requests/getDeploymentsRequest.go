package requests

import (
	"fmt"
	"github.com/aljrubior/anyctl/clients"
	"github.com/aljrubior/anyctl/conf"
	"net/http"
)

func NewGetDeploymentsRequest(
	config *conf.DeploymentClientConfig,
	bearerToken,
	organizationId,
	environmentId string) *GetDeploymentsRequest {

	return &GetDeploymentsRequest{
		config:         config,
		bearerToken:    bearerToken,
		organizationId: organizationId,
		environmentId:  environmentId,
	}
}

type GetDeploymentsRequest struct {
	clients.BaseHttpRequest
	config         *conf.DeploymentClientConfig
	bearerToken    string
	organizationId string
	environmentId  string
}

func (this *GetDeploymentsRequest) buildUri() string {

	protocol := this.config.Protocol
	host := this.config.Host
	path := fmt.Sprintf(this.config.DeploymentsPath, this.organizationId, this.environmentId)

	return fmt.Sprintf("%s://%s/%s", protocol, host, path)
}

func (this *GetDeploymentsRequest) Build() *http.Request {

	uri := this.buildUri()

	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	this.AddDefaultHeaders(req, this.organizationId, this.environmentId, this.bearerToken)

	return req
}
