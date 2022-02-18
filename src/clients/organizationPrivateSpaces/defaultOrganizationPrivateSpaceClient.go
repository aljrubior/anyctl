package organizationPrivateSpaces

import (
	"encoding/json"
	"github.com/aljrubior/anyctl/clients"
	"github.com/aljrubior/anyctl/clients/organizationPrivateSpaces/requests"
	"github.com/aljrubior/anyctl/clients/organizationPrivateSpaces/response"
	"github.com/aljrubior/anyctl/conf"
	"io/ioutil"
	"net/http"
	"time"
)

func NewOrganizationDefaultPrivateSpaceClient(config conf.RuntimeFabricClientConfig) DefaultOrganizationPrivateSpaceClient {
	return DefaultOrganizationPrivateSpaceClient{
		config: config,
	}
}

type DefaultOrganizationPrivateSpaceClient struct {
	clients.HttpClient
	config conf.RuntimeFabricClientConfig
}

func (this DefaultOrganizationPrivateSpaceClient) GetPrivateSpaces(orgId, envId, token string) (*response.OrganizationPrivateSpacesResponse, error) {

	client := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := requests.NewGetPrivateSpacesRequest(&this.config, token, orgId, envId).Build()

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

	var privatespaces response.OrganizationPrivateSpacesResponse

	err = json.Unmarshal(data, &privatespaces)

	if err != nil {
		return nil, err
	}

	return &privatespaces, nil
}

func (this DefaultOrganizationPrivateSpaceClient) GetPrivateSpace(orgId, envId, token, privateSpaceId string) (*response.OrganizationPrivateSpaceResponse, error) {

	client := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := requests.NewGetPrivateSpaceRequest(&this.config, token, orgId, envId, privateSpaceId).Build()

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

	var privatespace response.OrganizationPrivateSpaceResponse

	err = json.Unmarshal(data, &privatespace)

	if err != nil {
		return nil, err
	}

	return &privatespace, nil
}

func (this DefaultOrganizationPrivateSpaceClient) GetFabrics(orgId, envId, token, privateSpaceId string) (*[]response.OrganizationPrivateSpaceFabricResponse, error) {

	client := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := requests.NewGetPrivateSpaceFabricsRequest(&this.config, token, orgId, envId, privateSpaceId).Build()

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

	var fabrics []response.OrganizationPrivateSpaceFabricResponse

	err = json.Unmarshal(data, &fabrics)

	if err != nil {
		return nil, err
	}

	return &fabrics, nil
}
