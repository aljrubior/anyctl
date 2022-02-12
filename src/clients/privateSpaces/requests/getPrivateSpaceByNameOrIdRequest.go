package requests

import (
	"fmt"
	"github.com/aljrubior/anyctl/clients"
	"github.com/aljrubior/anyctl/conf"
	"net/http"
)

func NewGetPrivateSpaceByNameOrIdRequest(config *conf.PrivateSpaceClientConfig, bearerToken, privateSpaceId string) *GetPrivateSpaceByNameOrIdRequest {
	return &GetPrivateSpaceByNameOrIdRequest{
		config:         config,
		bearerToken:    bearerToken,
		privateSpaceId: privateSpaceId,
	}
}

type GetPrivateSpaceByNameOrIdRequest struct {
	clients.BaseHttpRequest
	config         *conf.PrivateSpaceClientConfig
	bearerToken    string
	privateSpaceId string
}

func (this *GetPrivateSpaceByNameOrIdRequest) buildUri() string {

	protocol := this.config.Protocol
	host := this.config.Host
	path := this.config.PrivateSpacesPath

	return fmt.Sprintf("%s://%s/%s", protocol, host, path)
}

func (this *GetPrivateSpaceByNameOrIdRequest) Build() *http.Request {

	uri := this.buildUri()

	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	q := req.URL.Query()
	q.Add("page", "0")
	q.Add("privateSpaceNameOrId", this.privateSpaceId)
	req.URL.RawQuery = q.Encode()

	this.AddAuthorizationHeader(req, this.bearerToken)

	return req
}
