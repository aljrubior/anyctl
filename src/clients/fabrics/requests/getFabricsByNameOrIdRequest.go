package requests

import (
	"fmt"
	"github.com/aljrubior/anyctl/clients"
	"github.com/aljrubior/anyctl/conf"
	"net/http"
)

func NewGetFabricsByNameOrIdRequest(config *conf.FabricClientConfig, bearerToken, fabricId string) *GetFabricsByNameOrIdRequest {
	return &GetFabricsByNameOrIdRequest{
		config:      config,
		bearerToken: bearerToken,
		fabricId:    fabricId,
	}
}

type GetFabricsByNameOrIdRequest struct {
	clients.BaseHttpRequest
	config      *conf.FabricClientConfig
	bearerToken string
	fabricId    string
}

func (this *GetFabricsByNameOrIdRequest) buildUri() string {

	protocol := this.config.Protocol
	host := this.config.Host
	path := this.config.FabricsPath

	return fmt.Sprintf("%s://%s/%s", protocol, host, path)
}

func (this *GetFabricsByNameOrIdRequest) Build() *http.Request {

	uri := this.buildUri()

	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	q := req.URL.Query()
	q.Add("page", "0")
	q.Add("size", "25")
	q.Add("fabricNameOrFabricId", this.fabricId)
	req.URL.RawQuery = q.Encode()

	this.AddAuthorizationHeader(req, this.bearerToken)

	return req
}
