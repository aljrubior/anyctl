package requests

import (
	"fmt"
	"github.com/aljrubior/anyctl/clients"
	"github.com/aljrubior/anyctl/conf"
	"net/http"
)

func NewGetFabricsRequest(config conf.FabricClientConfig, bearerToken string) *GetFabricsRequest {
	return &GetFabricsRequest{
		config:      config,
		bearerToken: bearerToken,
	}
}

type GetFabricsRequest struct {
	clients.BaseHttpRequest
	config         conf.FabricClientConfig
	bearerToken    string
	privateSpaceId string
}

func (this *GetFabricsRequest) buildUri() string {

	protocol := this.config.Protocol
	host := this.config.Host
	path := this.config.FabricsPath

	return fmt.Sprintf("%s://%s/%s", protocol, host, path)
}

func (this *GetFabricsRequest) Build() *http.Request {

	uri := this.buildUri()

	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	this.AddAuthorizationHeader(req, this.bearerToken)

	return req
}
