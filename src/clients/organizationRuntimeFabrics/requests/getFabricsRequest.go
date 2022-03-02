package requests

import (
	"fmt"
	"github.com/aljrubior/anyctl/clients"
	"github.com/aljrubior/anyctl/conf"
	"net/http"
)

func NewGetFabricsRequest(
	config *conf.RuntimeFabricClientConfig,
	bearerToken,
	organizationId,
	environmentId string) *GetFabricsRequest {

	return &GetFabricsRequest{
		config:         config,
		bearerToken:    bearerToken,
		organizationId: organizationId,
		environmentId:  environmentId,
	}
}

type GetFabricsRequest struct {
	clients.BaseHttpRequest
	config         *conf.RuntimeFabricClientConfig
	bearerToken    string
	organizationId string
	environmentId  string
}

func (this *GetFabricsRequest) buildUri() string {

	protocol := this.config.Protocol
	host := this.config.Host
	path := fmt.Sprintf(this.config.FabricsPath, this.organizationId)
	return fmt.Sprintf("%s://%s/%s", protocol, host, path)
}

func (this *GetFabricsRequest) Build() *http.Request {

	uri := this.buildUri()

	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	this.AddDefaultHeaders(req, this.organizationId, this.environmentId, this.bearerToken)

	return req
}
