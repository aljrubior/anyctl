package organizationRuntimeFabrics

import (
	"encoding/json"
	"github.com/aljrubior/anyctl/clients"
	"github.com/aljrubior/anyctl/clients/organizationRuntimeFabrics/requests"
	"github.com/aljrubior/anyctl/clients/organizationRuntimeFabrics/response"
	"github.com/aljrubior/anyctl/conf"
	"io/ioutil"
	"net/http"
	"time"
)

func NewDefaultOrganizationRuntimeFabricClient(config *conf.RuntimeFabricClientConfig) *DefaultOrganizationRuntimeFabricClient {
	return &DefaultOrganizationRuntimeFabricClient{
		config: config,
	}
}

type DefaultOrganizationRuntimeFabricClient struct {
	clients.HttpClient
	config *conf.RuntimeFabricClientConfig
}

func (this *DefaultOrganizationRuntimeFabricClient) GetFabrics(orgId, envId, token string) (*[]response.OrganizationFabricResponse, error) {

	client := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := requests.NewGetFabricsRequest(this.config, token, orgId, envId).Build()

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

	var targets []response.OrganizationFabricResponse

	err = json.Unmarshal(data, &targets)

	if err != nil {
		return nil, err
	}

	return &targets, nil
}

func (this *DefaultOrganizationRuntimeFabricClient) GetFabric(orgId, envId, token, targetId string) (*response.OrganizationFabricResponse, error) {

	client := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := requests.NewGetFabricRequest(this.config, token, orgId, envId, targetId).Build()

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

	var target response.OrganizationFabricResponse

	err = json.Unmarshal(data, &target)

	if err != nil {
		return nil, err
	}

	return &target, nil
}

func (this *DefaultOrganizationRuntimeFabricClient) GetTargets(orgId, envId, token string) (*[]response.FabricTargetResponse, error) {

	client := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := requests.NewGetTargetsRequest(this.config, token, orgId, envId).Build()

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

	var targets []response.FabricTargetResponse

	err = json.Unmarshal(data, &targets)

	if err != nil {
		return nil, err
	}

	return &targets, nil
}

func (this *DefaultOrganizationRuntimeFabricClient) GetTarget(orgId, envId, token, targetId string) (*response.FabricTargetResponse, error) {

	client := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := requests.NewGetTargetRequest(this.config, token, orgId, envId, targetId).Build()

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

	var target response.FabricTargetResponse

	err = json.Unmarshal(data, &target)

	if err != nil {
		return nil, err
	}

	return &target, nil
}
