package fabrics

import (
	"encoding/json"
	"github.com/aljrubior/anyctl/clients"
	"github.com/aljrubior/anyctl/clients/fabrics/requests"
	"github.com/aljrubior/anyctl/clients/fabrics/response"
	"github.com/aljrubior/anyctl/conf"
	"io/ioutil"
	"net/http"
	"time"
)

func NewDefaultFabricClient(config conf.FabricClientConfig) DefaultFabricClient {
	return DefaultFabricClient{
		config: config,
	}
}

type DefaultFabricClient struct {
	clients.HttpClient
	config conf.FabricClientConfig
}

func (this DefaultFabricClient) GetFabrics(token string) (*[]response.FabricResponse, error) {

	client := &http.Client{Timeout: time.Duration(10) * time.Second}

	println(token)

	req := requests.NewGetFabricsRequest(&this.config, token).Build()

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

	var fabrics []response.FabricResponse

	err = json.Unmarshal(data, &fabrics)

	if err != nil {
		return nil, err
	}

	return &fabrics, nil
}

func (this DefaultFabricClient) GetFabric(token, fabricId string) (*response.FabricResponse, error) {

	client := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := requests.NewGetFabricRequest(&this.config, token, fabricId).Build()

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

	var fabric response.FabricResponse

	err = json.Unmarshal(data, &fabric)

	if err != nil {
		return nil, err
	}

	return &fabric, nil
}

func (this DefaultFabricClient) GetFabricsByNameOrId(token string, fabricId string) (*response.FabricsResponse, error) {

	client := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := requests.NewGetFabricsByNameOrIdRequest(&this.config, token, fabricId).Build()

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

	var fabricsResponse response.FabricsResponse

	err = json.Unmarshal(data, &fabricsResponse)

	if err != nil {
		return nil, err
	}

	return &fabricsResponse, nil
}
