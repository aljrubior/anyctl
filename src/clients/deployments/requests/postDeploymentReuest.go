package requests

import (
	"bytes"
	"fmt"
	"github.com/aljrubior/anyctl/clients"
	"github.com/aljrubior/anyctl/conf"
	"net/http"
)

func NewPostDeploymentRequest(
	config *conf.DeploymentClientConfig,
	bearerToken,
	organizationId,
	environmentId string,
	body []byte) *PostDeploymentRequest {

	return &PostDeploymentRequest{
		config:         config,
		bearerToken:    bearerToken,
		organizationId: organizationId,
		environmentId:  environmentId,
		body:           body,
	}
}

type PostDeploymentRequest struct {
	clients.BaseHttpRequest
	config         *conf.DeploymentClientConfig
	bearerToken    string
	organizationId string
	environmentId  string
	body           []byte
}

func (this *PostDeploymentRequest) buildUri() string {

	protocol := this.config.Protocol
	host := this.config.Host
	path := fmt.Sprintf(this.config.DeploymentsPathTemplate, this.organizationId, this.environmentId)

	return fmt.Sprintf("%s://%s/%s", protocol, host, path)
}

func (this *PostDeploymentRequest) Build() *http.Request {

	uri := this.buildUri()

	req, _ := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(this.body))

	this.AddDefaultHeaders(req, this.organizationId, this.environmentId, this.bearerToken)

	this.AddContentType(req, clients.ContentTypeJSON)

	return req
}
