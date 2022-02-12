package schedulers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aljrubior/anyctl/clients"
	"github.com/aljrubior/anyctl/clients/schedulers/response"
	"github.com/aljrubior/anyctl/conf"
	"io/ioutil"
	"net/http"
	"time"
)

func NewDefaultSchedulerClient(config *conf.SchedulerClientConfig) *DefaultSchedulerClient {
	return &DefaultSchedulerClient{
		config: config,
	}

}

type DefaultSchedulerClient struct {
	clients.HttpClient
	config *conf.SchedulerClientConfig
}

func (this DefaultSchedulerClient) GetSchedulers(orgId, envId, token, deploymentId string) (*response.SchedulersResponse, error) {

	client := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := this.buildGetSchedulersRequest(orgId, envId, token, deploymentId)

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return nil, this.ThrowError(resp)
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

	req := this.buildPatchSchedulersRequest(orgId, envId, token, deploymentId, body)

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 202 {
		return nil, this.ThrowError(resp)
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

	req := this.buildPostSchedulerRequest(orgId, envId, token, deploymentId, flowName)

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

	req := this.buildDeleteSchedulerRequest(orgId, envId, token, deploymentId, flowName)

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

func (this DefaultSchedulerClient) buildGetSchedulersRequest(orgId, envId, token, deploymentId string) *http.Request {

	uri := this.buildSchedulersUri(orgId, envId, deploymentId)

	req, _ := http.NewRequest("GET", uri, nil)
	this.AddDefaultHeaders(req, orgId, envId, token)

	return req
}

func (this DefaultSchedulerClient) buildPatchSchedulersRequest(orgId, envId, token, deploymentId string, body []byte) *http.Request {

	uri := this.buildSchedulersUri(orgId, envId, deploymentId)

	req, _ := http.NewRequest("PATCH", uri, bytes.NewBuffer(body))

	this.AddDefaultHeaders(req, orgId, envId, token)
	this.AddContentTypeApplicationJsonHeader(req)

	return req
}

func (this DefaultSchedulerClient) buildPostSchedulerRequest(orgId, envId, token, deploymentId string, flowName string) *http.Request {

	uri := this.buildRunSchedulersUri(orgId, envId, deploymentId, flowName)

	req, _ := http.NewRequest("POST", uri, nil)

	this.AddDefaultHeaders(req, orgId, envId, token)

	return req
}

func (this DefaultSchedulerClient) buildDeleteSchedulerRequest(orgId, envId, token, deploymentId string, flowName string) *http.Request {

	uri := this.buildDeleteSchedulersUri(orgId, envId, deploymentId, flowName)

	req, _ := http.NewRequest(http.MethodDelete, uri, nil)

	this.AddDefaultHeaders(req, orgId, envId, token)

	return req
}

func (this DefaultSchedulerClient) buildSchedulersUri(orgId, envId, deploymentId string) string {

	protocol := this.config.Protocol
	host := this.config.Host
	path := this.config.SchedulersPathTemplate
	schedulersPath := fmt.Sprintf(path, orgId, envId, deploymentId)
	return fmt.Sprintf("%s://%s/%s", protocol, host, schedulersPath)
}

func (this DefaultSchedulerClient) buildRunSchedulersUri(orgId, envId, deploymentId, flowName string) string {

	protocol := this.config.Protocol
	host := this.config.Host
	path := this.config.RunSchedulerPathTemplate
	runSchedulerPath := fmt.Sprintf(path, orgId, envId, deploymentId, flowName)

	return fmt.Sprintf("%s://%s/%s", protocol, host, runSchedulerPath)
}

func (this DefaultSchedulerClient) buildDeleteSchedulersUri(orgId, envId, deploymentId, flowName string) string {

	protocol := this.config.Protocol
	host := this.config.Host
	path := fmt.Sprintf(this.config.SchedulerPathTemplate, orgId, envId, deploymentId, flowName)

	return fmt.Sprintf("%s://%s/%s", protocol, host, path)
}
