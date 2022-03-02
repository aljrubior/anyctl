package sharedspaces

import (
	"encoding/json"
	"github.com/aljrubior/anyctl/clients"
	"github.com/aljrubior/anyctl/clients/sharedspaces/requests"
	"github.com/aljrubior/anyctl/clients/sharedspaces/response"
	"github.com/aljrubior/anyctl/conf"
	"io/ioutil"
	"net/http"
	"time"
)

func NewDefaultSharedSpaceClient(config conf.SharedSpaceClientConfig) DefaultSharedSpaceClient {
	return DefaultSharedSpaceClient{
		config: config,
	}
}

type DefaultSharedSpaceClient struct {
	clients.HttpClient
	config conf.SharedSpaceClientConfig
}

func (this DefaultSharedSpaceClient) GetSharedSpace(token, sharedSpaceId string) (*response.SharedSpaceResponse, error) {

	client := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := requests.NewGetSharedSpaceRequest(&this.config, token, sharedSpaceId).Build()

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, this.ThrowError(resp)
	}

	var privatespace response.SharedSpaceResponse

	err = json.Unmarshal(data, &privatespace)

	if err != nil {
		return nil, err
	}

	return &privatespace, nil
}

func (this DefaultSharedSpaceClient) GetSharedSpaces(token string) (*[]response.SharedSpaceResponse, error) {

	client := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := requests.NewGetSharedSpacesRequest(&this.config, token).Build()

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, this.ThrowError(resp)
	}

	var sharedspaces []response.SharedSpaceResponse

	err = json.Unmarshal(data, &sharedspaces)

	if err != nil {
		return nil, err
	}

	return &sharedspaces, nil
}
