package accounts

import (
	"encoding/json"
	"github.com/aljrubior/anyctl/clients"
	"github.com/aljrubior/anyctl/clients/accounts/requests"
	"github.com/aljrubior/anyctl/clients/accounts/response"
	"github.com/aljrubior/anyctl/conf"
	"io/ioutil"
	"net/http"
	"time"
)

func NewDefaultAccountClient(config conf.AccountClientConfig) DefaultAccountClient {
	return DefaultAccountClient{
		config: config,
	}
}

type DefaultAccountClient struct {
	clients.HttpClient
	config conf.AccountClientConfig
}

func (this DefaultAccountClient) GetOrganization(token, orgId string) (*response.OrganizationResponse, error) {

	client := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := requests.NewGetOrganizationRequest(&this.config, token, orgId).Build()

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return nil, this.ThrowError(resp)
	}

	var response response.OrganizationResponse

	err = json.Unmarshal(data, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (this DefaultAccountClient) GetProfile(token string) (*response.Profile, error) {

	client := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := requests.NewGetProfileRequest(&this.config, token).Build()

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, this.ThrowError(resp)
	}

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var profile response.Profile

	err = json.Unmarshal(data, &profile)

	if err != nil {
		return nil, err
	}

	return &profile, nil
}

func (this DefaultAccountClient) Login(body []byte) (*response.LoginResponse, error) {

	client := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := requests.NewPostLoginRequest(&this.config, body).Build()

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, this.ThrowError(resp)
	}

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var loginResponse response.LoginResponse

	if err := json.Unmarshal(data, &loginResponse); err != nil {
		return nil, err
	}

	return &loginResponse, nil
}
