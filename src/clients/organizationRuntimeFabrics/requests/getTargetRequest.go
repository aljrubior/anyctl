package requests

import (
	"fmt"
	"github.com/aljrubior/anyctl/clients"
	"github.com/aljrubior/anyctl/conf"
	"net/http"
)

func NewGetTargetRequest(
	config *conf.RuntimeFabricClientConfig,
	bearerToken,
	organizationId,
	environmentId,
	targetId string) *GetTargetRequest {

	return &GetTargetRequest{
		config:         config,
		bearerToken:    bearerToken,
		organizationId: organizationId,
		environmentId:  environmentId,
		targetId:       targetId,
	}
}

type GetTargetRequest struct {
	clients.BaseHttpRequest
	config         *conf.RuntimeFabricClientConfig
	bearerToken    string
	organizationId string
	environmentId  string
	targetId       string
}

func (this *GetTargetRequest) buildUri() string {

	protocol := this.config.Protocol
	host := this.config.Host
	path := fmt.Sprintf(this.config.TargetPath, this.organizationId, this.targetId)
	return fmt.Sprintf("%s://%s/%s", protocol, host, path)
}

func (this *GetTargetRequest) Build() *http.Request {

	uri := this.buildUri()

	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	this.AddDefaultHeaders(req, this.organizationId, this.environmentId, this.bearerToken)

	return req
}
