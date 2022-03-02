package requests

import (
	"fmt"
	"github.com/aljrubior/anyctl/clients"
	"github.com/aljrubior/anyctl/conf"
	"net/http"
)

func NewGetSharedSpacesRequest(config *conf.SharedSpaceClientConfig, bearerToken string) *GetSharedSpacesRequest {
	return &GetSharedSpacesRequest{
		config:      config,
		bearerToken: bearerToken,
	}
}

type GetSharedSpacesRequest struct {
	clients.BaseHttpRequest
	config      *conf.SharedSpaceClientConfig
	bearerToken string
}

func (this *GetSharedSpacesRequest) buildUri() string {

	protocol := this.config.Protocol
	host := this.config.Host
	path := this.config.SharedSpacesPath

	return fmt.Sprintf("%s://%s/%s", protocol, host, path)
}

func (this *GetSharedSpacesRequest) Build() *http.Request {

	uri := this.buildUri()

	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	this.AddAuthorizationHeader(req, this.bearerToken)

	return req
}
