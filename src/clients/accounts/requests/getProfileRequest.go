package requests

import (
	"fmt"
	"github.com/aljrubior/anyctl/clients"
	"github.com/aljrubior/anyctl/conf"
	"net/http"
)

func NewGetProfileRequest(config conf.AccountClientConfig, bearerToken string) *GetProfileRequest {
	return &GetProfileRequest{
		config:      config,
		bearerToken: bearerToken,
	}
}

type GetProfileRequest struct {
	clients.BaseHttpRequest
	config      conf.AccountClientConfig
	bearerToken string
}

func (this *GetProfileRequest) buildUri() string {

	protocol := this.config.Protocol
	host := this.config.Host
	path := this.config.ProfilePath
	return fmt.Sprintf("%s://%s/%s", protocol, host, path)
}

func (this *GetProfileRequest) Build() *http.Request {

	uri := this.buildUri()

	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	this.AddAuthorizationHeader(req, this.bearerToken)

	return req
}
