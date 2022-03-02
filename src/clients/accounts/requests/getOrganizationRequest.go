package requests

import (
	"fmt"
	"github.com/aljrubior/anyctl/clients"
	"github.com/aljrubior/anyctl/conf"
	"net/http"
)

func NewGetOrganizationRequest(config *conf.AccountClientConfig, bearerToken, organizationId string) *GetOrganizationRequest {
	return &GetOrganizationRequest{
		config:         config,
		bearerToken:    bearerToken,
		organizationId: organizationId,
	}
}

type GetOrganizationRequest struct {
	clients.BaseHttpRequest
	config         *conf.AccountClientConfig
	bearerToken    string
	organizationId string
}

func (this *GetOrganizationRequest) buildUri() string {

	protocol := this.config.Protocol
	host := this.config.Host
	path := this.config.OrganizationPath
	organizationPath := fmt.Sprintf(path, this.organizationId)
	return fmt.Sprintf("%s://%s/%s", protocol, host, organizationPath)
}

func (this *GetOrganizationRequest) Build() *http.Request {

	uri := this.buildUri()

	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	this.AddAuthorizationHeader(req, this.bearerToken)

	return req
}
