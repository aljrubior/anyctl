package requests

import (
	"fmt"
	"github.com/aljrubior/anyctl/clients"
	"github.com/aljrubior/anyctl/conf"
	"net/http"
)

func NewGetPrivateSpacesRequest(config *conf.PrivateSpaceClientConfig, bearerToken string) *GetPrivateSpacesRequest {
	return &GetPrivateSpacesRequest{
		config:      config,
		bearerToken: bearerToken,
	}
}

type GetPrivateSpacesRequest struct {
	clients.BaseHttpRequest
	config      *conf.PrivateSpaceClientConfig
	bearerToken string
}

func (this *GetPrivateSpacesRequest) buildUri() string {

	protocol := this.config.Protocol
	host := this.config.Host
	path := this.config.PrivateSpacesPath

	return fmt.Sprintf("%s://%s/%s", protocol, host, path)
}

func (this *GetPrivateSpacesRequest) Build() *http.Request {

	uri := this.buildUri()

	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	this.AddAuthorizationHeader(req, this.bearerToken)

	return req
}
