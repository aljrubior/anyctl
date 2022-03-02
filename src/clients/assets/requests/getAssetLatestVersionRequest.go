package requests

import (
	"fmt"
	"github.com/aljrubior/anyctl/clients"
	"github.com/aljrubior/anyctl/conf"
	"net/http"
)

func NewGetAssetLatestVersionRequest(
	config *conf.AssetClientConfig,
	bearerToken,
	organizationId,
	assetName string) *GetAssetLatestVersionRequest {

	return &GetAssetLatestVersionRequest{
		config:         config,
		bearerToken:    bearerToken,
		organizationId: organizationId,
		assetName:      assetName,
	}
}

type GetAssetLatestVersionRequest struct {
	clients.BaseHttpRequest
	config         *conf.AssetClientConfig
	bearerToken    string
	organizationId string
	assetName      string
}

func (this *GetAssetLatestVersionRequest) buildUri() string {

	protocol := this.config.Protocol
	host := this.config.Host
	path := fmt.Sprintf(this.config.LatestVersionPath, this.organizationId, this.assetName)

	return fmt.Sprintf("%s://%s/%s", protocol, host, path)
}

func (this *GetAssetLatestVersionRequest) Build() *http.Request {

	uri := this.buildUri()

	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	this.AddAuthorizationHeader(req, this.bearerToken)

	return req
}
