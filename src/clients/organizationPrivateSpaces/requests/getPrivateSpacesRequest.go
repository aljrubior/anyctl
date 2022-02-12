package requests

import (
	"fmt"
	"github.com/aljrubior/anyctl/clients"
	"github.com/aljrubior/anyctl/conf"
	"net/http"
)

func NewGetPrivateSpacesRequest(
	config conf.RuntimeFabricClientConfig,
	bearerToken,
	organizationId,
	environmentId string) *GetPrivateSpacesRequest {

	return &GetPrivateSpacesRequest{
		config:         config,
		bearerToken:    bearerToken,
		organizationId: organizationId,
		environmentId:  environmentId,
	}
}

type GetPrivateSpacesRequest struct {
	clients.BaseHttpRequest
	config         conf.RuntimeFabricClientConfig
	bearerToken    string
	organizationId string
	environmentId  string
}

func (this *GetPrivateSpacesRequest) buildUri() string {

	protocol := this.config.Protocol
	host := this.config.Host
	path := fmt.Sprintf(this.config.PrivateSpacesPath, this.organizationId)

	return fmt.Sprintf("%s://%s/%s", protocol, host, path)
}

func (this *GetPrivateSpacesRequest) Build() *http.Request {

	uri := this.buildUri()

	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	this.AddDefaultHeaders(req, this.organizationId, this.environmentId, this.bearerToken)

	return req
}
