package requests

import (
	"fmt"
	"github.com/aljrubior/anyctl/clients"
	"github.com/aljrubior/anyctl/conf"
	"net/http"
)

func NewGetFabricsRequest(config *conf.PrivateSpaceClientConfig, bearerToken, privateSpaceId string) *GetFabricsRequest {
	return &GetFabricsRequest{
		config:         config,
		bearerToken:    bearerToken,
		privateSpaceId: privateSpaceId,
	}
}

type GetFabricsRequest struct {
	clients.BaseHttpRequest
	config         *conf.PrivateSpaceClientConfig
	bearerToken    string
	privateSpaceId string
}

func (this *GetFabricsRequest) buildUri() string {

	protocol := this.config.Protocol
	host := this.config.Host
	path := fmt.Sprintf(this.config.FabricsPath, this.privateSpaceId)

	return fmt.Sprintf("%s://%s/%s", protocol, host, path)
}

func (this *GetFabricsRequest) Build() *http.Request {

	uri := this.buildUri()

	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	this.AddAuthorizationHeader(req, this.bearerToken)

	return req
}
