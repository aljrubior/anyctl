package requests

import (
	"fmt"
	"github.com/aljrubior/anyctl/clients"
	"github.com/aljrubior/anyctl/conf"
	"net/http"
)

func NewGetAssetVersionRequest(
	config *conf.AssetClientConfig,
	bearerToken,
	organizationId,
	assetName,
	assetVersion string) *GetAssetVersionRequest {

	return &GetAssetVersionRequest{
		config:         config,
		bearerToken:    bearerToken,
		organizationId: organizationId,
		assetName:      assetName,
		assetVersion:   assetVersion,
	}
}

type GetAssetVersionRequest struct {
	clients.BaseHttpRequest
	config         *conf.AssetClientConfig
	bearerToken    string
	organizationId string
	assetName      string
	assetVersion   string
}

func (this *GetAssetVersionRequest) buildUri() string {

	protocol := this.config.Protocol
	host := this.config.Host
	path := fmt.Sprintf(this.config.SpecificVersionPath, this.organizationId, this.assetName, this.assetVersion)

	return fmt.Sprintf("%s://%s/%s", protocol, host, path)
}

func (this *GetAssetVersionRequest) Build() *http.Request {

	uri := this.buildUri()

	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	this.AddAuthorizationHeader(req, this.bearerToken)

	return req
}
