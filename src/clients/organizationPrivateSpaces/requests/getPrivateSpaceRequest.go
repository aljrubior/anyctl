package requests

import (
	"fmt"
	"github.com/aljrubior/anyctl/clients"
	"github.com/aljrubior/anyctl/conf"
	"net/http"
)

func NewGetPrivateSpaceRequest(
	config *conf.RuntimeFabricClientConfig,
	bearerToken,
	organizationId,
	environmentId,
	privateSpaceId string) *GetPrivateSpaceRequest {

	return &GetPrivateSpaceRequest{
		config:         config,
		bearerToken:    bearerToken,
		organizationId: organizationId,
		environmentId:  environmentId,
		privateSpaceId: privateSpaceId,
	}
}

type GetPrivateSpaceRequest struct {
	clients.BaseHttpRequest
	config         *conf.RuntimeFabricClientConfig
	bearerToken    string
	organizationId string
	environmentId  string
	privateSpaceId string
}

func (this *GetPrivateSpaceRequest) buildUri() string {

	protocol := this.config.Protocol
	host := this.config.Host
	path := fmt.Sprintf(this.config.PrivateSpacePath, this.organizationId, this.privateSpaceId)

	return fmt.Sprintf("%s://%s/%s", protocol, host, path)
}

func (this *GetPrivateSpaceRequest) Build() *http.Request {

	uri := this.buildUri()

	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	this.AddDefaultHeaders(req, this.organizationId, this.environmentId, this.bearerToken)

	return req
}
