package requests

import (
	"fmt"
	"github.com/aljrubior/anyctl/clients"
	"github.com/aljrubior/anyctl/conf"
	"net/http"
)

func NewGetPrivateSpaceFabricsRequest(
	config conf.RuntimeFabricClientConfig,
	bearerToken,
	organizationId,
	environmentId,
	privateSpaceId string) *GetPrivateSpaceFabricsRequest {

	return &GetPrivateSpaceFabricsRequest{
		config:         config,
		bearerToken:    bearerToken,
		organizationId: organizationId,
		environmentId:  environmentId,
		privateSpaceId: privateSpaceId,
	}
}

type GetPrivateSpaceFabricsRequest struct {
	clients.BaseHttpRequest
	config         conf.RuntimeFabricClientConfig
	bearerToken    string
	organizationId string
	environmentId  string
	privateSpaceId string
}

func (this *GetPrivateSpaceFabricsRequest) buildUri() string {

	protocol := this.config.Protocol
	host := this.config.Host
	path := fmt.Sprintf(this.config.PrivateSpaceFabricPath, this.organizationId, this.privateSpaceId)

	return fmt.Sprintf("%s://%s/%s", protocol, host, path)
}

func (this *GetPrivateSpaceFabricsRequest) Build() *http.Request {

	uri := this.buildUri()

	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	this.AddDefaultHeaders(req, this.organizationId, this.environmentId, this.bearerToken)

	return req
}
