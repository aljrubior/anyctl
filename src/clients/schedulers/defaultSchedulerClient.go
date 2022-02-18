package schedulers

import (
	"encoding/json"
	"github.com/aljrubior/anyctl/clients"
	"github.com/aljrubior/anyctl/clients/schedulers/requests"
	"github.com/aljrubior/anyctl/clients/schedulers/response"
	"github.com/aljrubior/anyctl/conf"
	"io/ioutil"
	"net/http"
	"time"
)

func NewDefaultSchedulerClient(config conf.SchedulerClientConfig) DefaultSchedulerClient {

	return DefaultSchedulerClient{
		config: config,
	}
}

type DefaultSchedulerClient struct {
	clients.HttpClient
	config conf.SchedulerClientConfig
}

func (this DefaultSchedulerClient) GetSchedulers(orgId, envId, token, deploymentId string) (*response.SchedulersResponse, error) {

	client := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := requests.NewGetSchedulersRequest(&this.config, token, orgId, envId, deploymentId).Build()

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

	var response response.SchedulersResponse

	err = json.Unmarshal(data, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (this DefaultSchedulerClient) PatchSchedulers(orgId, envId, token, deploymentId string, body []byte) (*response.SchedulersResponse, error) {

	client := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := requests.NewPatchSchedulersRequest(&this.config, token, orgId, envId, deploymentId, body).Build()

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 202 {
		return nil, this.ThrowError(resp)
	}

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var response response.SchedulersResponse

	err = json.Unmarshal(data, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (this DefaultSchedulerClient) PostScheduler(orgId, envId, token, deploymentId string, flowName string) error {

	client := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := requests.NewPostSchedulerRequest(&this.config, token, orgId, envId, deploymentId, flowName).Build()

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return err
	}

	if resp.StatusCode != 204 {
		return this.ThrowError(resp)
	}

	return nil
}

func (this DefaultSchedulerClient) DeleteScheduler(orgId, envId, token, deploymentId string, flowName string) error {

	client := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := requests.NewDeleteSchedulerRequest(&this.config, token, orgId, envId, deploymentId, flowName).Build()

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return this.ThrowError(resp)
	}

	return nil
}
