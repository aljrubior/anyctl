package requests

import (
	"fmt"
	"github.com/aljrubior/anyctl/clients"
	"github.com/aljrubior/anyctl/conf"
	"net/http"
)

func NewGetFabricRequest(config *conf.PrivateSpaceClientConfig, bearerToken, privateSpaceId, fabricId string) *GetFabricRequest {
	return &GetFabricRequest{
		config:         config,
		bearerToken:    bearerToken,
		privateSpaceId: privateSpaceId,
		fabricId:       fabricId,
	}
}

type GetFabricRequest struct {
	clients.BaseHttpRequest
	config         *conf.PrivateSpaceClientConfig
	bearerToken    string
	privateSpaceId string
	fabricId       string
}

func (this *GetFabricRequest) buildUri() string {

	protocol := this.config.Protocol
	host := this.config.Host
	path := fmt.Sprintf(this.config.FabricsPath, this.privateSpaceId, this.fabricId)

	return fmt.Sprintf("%s://%s/%s", protocol, host, path)
}

func (this *GetFabricRequest) Build() *http.Request {

	uri := this.buildUri()

	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	this.AddAuthorizationHeader(req, this.bearerToken)

	return req
}
