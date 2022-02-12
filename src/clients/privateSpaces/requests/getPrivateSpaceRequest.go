package requests

import (
	"fmt"
	"github.com/aljrubior/anyctl/clients"
	"github.com/aljrubior/anyctl/conf"
	"net/http"
)

func NewGetPrivateSpaceRequest(config *conf.PrivateSpaceClientConfig, bearerToken, privateSpaceId string) *GetPrivateSpaceRequest {
	return &GetPrivateSpaceRequest{
		config:         config,
		bearerToken:    bearerToken,
		privateSpaceId: privateSpaceId,
	}
}

type GetPrivateSpaceRequest struct {
	clients.BaseHttpRequest
	config         *conf.PrivateSpaceClientConfig
	bearerToken    string
	privateSpaceId string
}

func (this *GetPrivateSpaceRequest) buildUri() string {

	protocol := this.config.Protocol
	host := this.config.Host
	path := fmt.Sprintf(this.config.PrivateSpacePath, this.privateSpaceId)

	return fmt.Sprintf("%s://%s/%s", protocol, host, path)
}

func (this *GetPrivateSpaceRequest) Build() *http.Request {

	uri := this.buildUri()

	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	this.AddAuthorizationHeader(req, this.bearerToken)

	return req
}
