package deployments

import (
	"encoding/json"
	"github.com/aljrubior/anyctl/clients"
	"github.com/aljrubior/anyctl/clients/deployments/requests"
	"github.com/aljrubior/anyctl/clients/deployments/response"
	"github.com/aljrubior/anyctl/conf"
	"io/ioutil"
	"net/http"
	"time"
)

func NewDefaultDeploymentClient(config conf.DeploymentClientConfig) DefaultDeploymentClient {
	return DefaultDeploymentClient{
		config: config,
	}
}

type DefaultDeploymentClient struct {
	clients.HttpClient
	config conf.DeploymentClientConfig
}

func (this DefaultDeploymentClient) GetDeployments(orgId, envId, token string) (*response.DeploymentsResponse, error) {
	client := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := requests.NewGetDeploymentsRequest(&this.config, token, orgId, envId).Build()

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

	var response response.DeploymentsResponse

	err = json.Unmarshal(data, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (this DefaultDeploymentClient) GetDeployment(orgId, envId, token, deploymentId string) (*response.DeploymentResponse, error) {
	client := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := requests.NewGetDeploymentRequest(&this.config, token, orgId, envId, deploymentId).Build()

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

	var response response.DeploymentResponse

	err = json.Unmarshal(data, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (this DefaultDeploymentClient) PostDeployment(orgId, envId, token string, body []byte) (*response.DeploymentResponse, error) {
	client := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := requests.NewPostDeploymentRequest(&this.config, token, orgId, envId, body).Build()

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 202 {
		return nil, this.ThrowError(resp)
	}

	data, err := ioutil.ReadAll(resp.Body)

	var response response.DeploymentResponse

	err = json.Unmarshal(data, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (this DefaultDeploymentClient) PatchDeployment(orgId, envId, token, deploymentId string, body []byte) (*response.DeploymentResponse, error) {

	client := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := requests.NewPatchDeploymentRequest(&this.config, token, orgId, envId, deploymentId, body).Build()

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

	var response response.DeploymentResponse

	err = json.Unmarshal(data, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (this DefaultDeploymentClient) DeleteDeployment(orgId, envId, token, deploymentId string) error {

	client := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := requests.NewDeleteDeploymentRequest(&this.config, token, orgId, envId, deploymentId).Build()

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

func (this DefaultDeploymentClient) GetDeploymentSpecs(orgId, envId, token, deploymentId string) (*[]response.DeploymentSpecResponse, error) {
	client := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := requests.NewGetDeploymentSpecsRequest(&this.config, token, orgId, envId, deploymentId).Build()

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return nil, this.ThrowError(resp)
	}

	var response []response.DeploymentSpecResponse

	err = json.Unmarshal(data, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}
