package requests

import (
	"fmt"
	"github.com/aljrubior/anyctl/clients"
	"github.com/aljrubior/anyctl/conf"
	"net/http"
)

func NewGetSharedSpaceRequest(config *conf.SharedSpaceClientConfig, bearerToken, sharedSpaceId string) *GetSharedSpaceRequest {
	return &GetSharedSpaceRequest{
		config:        config,
		bearerToken:   bearerToken,
		sharedSpaceId: sharedSpaceId,
	}
}

type GetSharedSpaceRequest struct {
	clients.BaseHttpRequest
	config        *conf.SharedSpaceClientConfig
	bearerToken   string
	sharedSpaceId string
}

func (this *GetSharedSpaceRequest) buildUri() string {

	protocol := this.config.Protocol
	host := this.config.Host
	path := fmt.Sprintf(this.config.SharedSpacePath, this.sharedSpaceId)

	return fmt.Sprintf("%s://%s/%s", protocol, host, path)
}

func (this *GetSharedSpaceRequest) Build() *http.Request {

	uri := this.buildUri()

	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	this.AddAuthorizationHeader(req, this.bearerToken)

	return req
}
