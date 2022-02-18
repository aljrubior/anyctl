package requests

import (
	"bytes"
	"fmt"
	"github.com/aljrubior/anyctl/clients"
	"github.com/aljrubior/anyctl/conf"
	"net/http"
)

func NewPostLoginRequest(config *conf.AccountClientConfig, body []byte) *PostLoginRequest {
	return &PostLoginRequest{
		config: config,
		body:   body,
	}
}

type PostLoginRequest struct {
	clients.BaseHttpRequest
	config *conf.AccountClientConfig
	body   []byte
}

func (this *PostLoginRequest) buildUri() string {

	protocol := this.config.Protocol
	host := this.config.Host
	path := this.config.LoginPath
	return fmt.Sprintf("%s://%s/%s", protocol, host, path)
}

func (this *PostLoginRequest) Build() *http.Request {

	uri := this.buildUri()

	req, _ := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(this.body))

	this.AddContentType(req, clients.ContentTypeJSON)

	return req
}
