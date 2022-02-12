package requests

import (
	"fmt"
	"github.com/aljrubior/anyctl/clients"
	"github.com/aljrubior/anyctl/conf"
	"net/http"
)

func NewGetFabricRequest(
	config *conf.RuntimeFabricClientConfig,
	bearerToken,
	organizationId,
	environmentId,
	targetId string) *GetFabricRequest {

	return &GetFabricRequest{
		config:         config,
		bearerToken:    bearerToken,
		organizationId: organizationId,
		environmentId:  environmentId,
		targetId:       targetId,
	}
}

type GetFabricRequest struct {
	clients.BaseHttpRequest
	config         *conf.RuntimeFabricClientConfig
	bearerToken    string
	organizationId string
	environmentId  string
	targetId       string
}

func (this *GetFabricRequest) buildUri() string {

	protocol := this.config.Protocol
	host := this.config.Host
	path := fmt.Sprintf(this.config.FabricPath, this.organizationId, this.targetId)
	return fmt.Sprintf("%s://%s/%s", protocol, host, path)
}

func (this *GetFabricRequest) Build() *http.Request {

	uri := this.buildUri()

	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	this.AddDefaultHeaders(req, this.organizationId, this.environmentId, this.bearerToken)

	return req
}
