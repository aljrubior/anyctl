package requests

import (
	"fmt"
	"github.com/aljrubior/anyctl/clients"
	"github.com/aljrubior/anyctl/conf"
	"net/http"
)

func NewGetTargets(
	config *conf.TargetClientConfig,
	organizationId,
	environmentId,
	bearerToken string) *GetTargets {

	return &GetTargets{
		config:         config,
		bearerToken:    bearerToken,
		organizationId: organizationId,
		environmentId:  environmentId,
	}
}

type GetTargets struct {
	clients.BaseHttpRequest
	config         *conf.TargetClientConfig
	bearerToken    string
	organizationId string
	environmentId  string
}

func (this *GetTargets) buildUri() string {

	protocol := this.config.Protocol
	host := this.config.Host
	path := this.config.TargetsPath

	return fmt.Sprintf("%s://%s/%s", protocol, host, path)
}

func (this *GetTargets) Build() *http.Request {

	uri := this.buildUri()

	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	this.AddDefaultHeaders(req, this.organizationId, this.environmentId, this.bearerToken)

	return req
}
