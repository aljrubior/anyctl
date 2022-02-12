package requests

import (
	"fmt"
	"github.com/aljrubior/anyctl/clients"
	"github.com/aljrubior/anyctl/conf"
	"net/http"
)

func NewPostLoginRequest(config conf.AccountClientConfig, username, password string) *PostLoginRequest {
	return &PostLoginRequest{
		config:   config,
		username: username,
		password: password,
	}
}

type PostLoginRequest struct {
	clients.BaseHttpRequest
	config   conf.AccountClientConfig
	username string
	password string
}

func (this *PostLoginRequest) buildUri() string {

	protocol := this.config.Protocol
	host := this.config.Host
	path := this.config.LoginPath
	return fmt.Sprintf("%s://%s/%s", protocol, host, path)
}

func (this *PostLoginRequest) Build() *http.Request {

	uri := this.buildUri()

	req, _ := http.NewRequest(http.MethodPost, uri, nil)
	param := req.URL.Query()
	param.Add("username", this.username)
	param.Add("password", this.password)

	req.URL.RawQuery = param.Encode()
	return req
}
