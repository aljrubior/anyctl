package requests

import (
	"bytes"
	"fmt"
	"github.com/aljrubior/anyctl/clients"
	"github.com/aljrubior/anyctl/conf"
	"net/http"
)

func NewPatchSchedulersRequest(
	config *conf.SchedulerClientConfig,
	bearerToken,
	organizationId,
	environmentId,
	deploymentId string,
	body []byte) *PatchSchedulersRequest {

	return &PatchSchedulersRequest{
		config:         config,
		bearerToken:    bearerToken,
		organizationId: organizationId,
		environmentId:  environmentId,
		deploymentId:   deploymentId,
		body:           body,
	}
}

type PatchSchedulersRequest struct {
	clients.BaseHttpRequest
	config         *conf.SchedulerClientConfig
	bearerToken    string
	organizationId string
	environmentId  string
	deploymentId   string
	body           []byte
}

func (this *PatchSchedulersRequest) buildUri() string {

	protocol := this.config.Protocol
	host := this.config.Host
	path := fmt.Sprintf(this.config.SchedulersPath, this.organizationId, this.environmentId, this.deploymentId)

	return fmt.Sprintf("%s://%s/%s", protocol, host, path)
}

func (this *PatchSchedulersRequest) Build() *http.Request {

	uri := this.buildUri()

	req, _ := http.NewRequest(http.MethodPatch, uri, bytes.NewBuffer(this.body))

	this.AddDefaultHeaders(req, this.organizationId, this.environmentId, this.bearerToken)
	this.AddContentType(req, clients.ContentTypeJSON)

	return req
}
