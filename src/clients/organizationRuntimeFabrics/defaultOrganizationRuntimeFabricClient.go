package organizationRuntimeFabrics

import (
	"encoding/json"
	"fmt"
	"github.com/aljrubior/anyctl/clients"
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

	req := this.buildFabricsRequest(orgId, envId, token)

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

	req := this.buildFabricRequest(orgId, envId, token, targetId)

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

	req := this.buildTargetsRequest(orgId, envId, token)

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

	req := this.buildTargetRequest(orgId, envId, token, targetId)

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

func (this *DefaultOrganizationRuntimeFabricClient) buildFabricsUri(orgId string) string {

	protocol := this.config.Protocol
	host := this.config.Host
	path := fmt.Sprintf(this.config.FabricsPath, orgId)
	return fmt.Sprintf("%s://%s/%s", protocol, host, path)
}

func (this *DefaultOrganizationRuntimeFabricClient) buildFabricUri(orgId, targetId string) string {

	protocol := this.config.Protocol
	host := this.config.Host
	path := fmt.Sprintf(this.config.FabricPath, orgId, targetId)
	return fmt.Sprintf("%s://%s/%s", protocol, host, path)
}

func (this *DefaultOrganizationRuntimeFabricClient) buildFabricsRequest(orgId, envId, token string) *http.Request {

	uri := this.buildFabricsUri(orgId)

	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	this.AddDefaultHeaders(req, orgId, envId, token)

	return req
}

func (this *DefaultOrganizationRuntimeFabricClient) buildFabricRequest(orgId, envId, token, targetId string) *http.Request {

	uri := this.buildFabricUri(orgId, targetId)

	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	this.AddDefaultHeaders(req, orgId, envId, token)

	return req
}

func (this *DefaultOrganizationRuntimeFabricClient) buildTargetUri(orgId, targetId string) string {

	protocol := this.config.Protocol
	host := this.config.Host
	path := fmt.Sprintf(this.config.TargetPath, orgId, targetId)
	return fmt.Sprintf("%s://%s/%s", protocol, host, path)
}

func (this *DefaultOrganizationRuntimeFabricClient) buildTargetRequest(orgId, envId, token, targetId string) *http.Request {

	uri := this.buildTargetUri(orgId, targetId)

	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	this.AddDefaultHeaders(req, orgId, envId, token)

	return req
}

func (this *DefaultOrganizationRuntimeFabricClient) buildTargetsUri(orgId string) string {

	protocol := this.config.Protocol
	host := this.config.Host
	path := fmt.Sprintf(this.config.TargetsPath, orgId)
	return fmt.Sprintf("%s://%s/%s", protocol, host, path)
}

func (this *DefaultOrganizationRuntimeFabricClient) buildTargetsRequest(orgId, envId, token string) *http.Request {

	uri := this.buildTargetsUri(orgId)

	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	this.AddDefaultHeaders(req, orgId, envId, token)

	return req
}
