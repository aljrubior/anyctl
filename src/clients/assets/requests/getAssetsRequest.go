package requests

import (
	"fmt"
	"github.com/aljrubior/anyctl/clients"
	"github.com/aljrubior/anyctl/conf"
	"net/http"
)

func NewGetAssetsRequest(
	config conf.AssetClientConfig,
	bearerToken,
	organizationId,
	environmentId,
	assetName string) *GetAssetsRequest {

	return &GetAssetsRequest{
		config:         config,
		bearerToken:    bearerToken,
		organizationId: organizationId,
		environmentId:  environmentId,
		assetName:      assetName,
	}
}

type GetAssetsRequest struct {
	clients.BaseHttpRequest
	config         conf.AssetClientConfig
	bearerToken    string
	organizationId string
	environmentId  string
	assetName      string
}

func (this *GetAssetsRequest) buildUri() string {

	protocol := this.config.Protocol
	host := this.config.Host
	path := this.config.AssetsPath
	return fmt.Sprintf("%s://%s/%s", protocol, host, path)
}

func (this *GetAssetsRequest) Build() *http.Request {

	uri := this.buildUri()

	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	this.AddAuthorizationHeader(req, this.bearerToken)
	this.AddDefaultHeaders(req, this.organizationId, this.environmentId, this.bearerToken)

	q := req.URL.Query()
	q.Add("search", this.assetName)
	q.Add("type", "app")
	req.URL.RawQuery = q.Encode()

	return req
}
