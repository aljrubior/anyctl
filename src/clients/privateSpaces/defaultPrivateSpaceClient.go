package privateSpaces

import (
	"encoding/json"
	"github.com/aljrubior/anyctl/clients"
	response2 "github.com/aljrubior/anyctl/clients/fabrics/response"
	"github.com/aljrubior/anyctl/clients/privateSpaces/requests"
	"github.com/aljrubior/anyctl/clients/privateSpaces/response"
	"github.com/aljrubior/anyctl/conf"
	"io/ioutil"
	"net/http"
	"time"
)

func NewDefaultPrivateSpaceClient(config conf.PrivateSpaceClientConfig) DefaultPrivateSpaceClient {
	return DefaultPrivateSpaceClient{
		config: config,
	}
}

type DefaultPrivateSpaceClient struct {
	clients.HttpClient
	config conf.PrivateSpaceClientConfig
}

func (this DefaultPrivateSpaceClient) GetPrivateSpaces(token string) (*response.PrivateSpacesResponse, error) {

	client := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := requests.NewGetPrivateSpacesRequest(&this.config, token).Build()

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

	var privateSpacesResponse response.PrivateSpacesResponse

	err = json.Unmarshal(data, &privateSpacesResponse)

	if err != nil {
		return nil, err
	}

	return &privateSpacesResponse, nil
}

func (this DefaultPrivateSpaceClient) GetPrivateSpace(token, privateSpaceId string) (*response.PrivateSpaceResponse, error) {

	client := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := requests.NewGetPrivateSpaceRequest(&this.config, token, privateSpaceId).Build()

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

	var privateSpaceResponse response.PrivateSpaceResponse

	err = json.Unmarshal(data, &privateSpaceResponse)

	if err != nil {
		return nil, err
	}

	return &privateSpaceResponse, nil
}

func (this DefaultPrivateSpaceClient) GetPrivateSpacesByNameOrId(token string, privateSpaceId string) (*response.PrivateSpacesResponse, error) {

	client := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := requests.NewGetPrivateSpaceByNameOrIdRequest(&this.config, token, privateSpaceId).Build()

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

	var privateSpacesResponse response.PrivateSpacesResponse

	err = json.Unmarshal(data, &privateSpacesResponse)

	if err != nil {
		return nil, err
	}

	return &privateSpacesResponse, nil
}

func (this DefaultPrivateSpaceClient) GetFabrics(token, privateSpaceId string) (*[]response2.FabricResponse, error) {

	client := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := requests.NewGetFabricsRequest(&this.config, token, privateSpaceId).Build()

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

	var fabrics []response2.FabricResponse

	err = json.Unmarshal(data, &fabrics)

	if err != nil {
		return nil, err
	}

	return &fabrics, nil
}

func (this DefaultPrivateSpaceClient) GetFabric(token, privateSpaceId, fabricId string) (*response2.FabricResponse, error) {

	client := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := requests.NewGetFabricRequest(&this.config, token, privateSpaceId, fabricId).Build()

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

	var fabric response2.FabricResponse

	err = json.Unmarshal(data, &fabric)

	if err != nil {
		return nil, err
	}

	return &fabric, nil
}
