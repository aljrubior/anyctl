package services

import (
	"github.com/aljrubior/anyctl/clients/organizationRuntimeFabrics"
	"github.com/aljrubior/anyctl/clients/organizationRuntimeFabrics/response"
)

func NewDefaultOrganizationRuntimeFabricService(runtimeFabricClient organizationRuntimeFabrics.OrganizationRuntimeFabricClient) DefaultOrganizationRuntimeFabricService {
	return DefaultOrganizationRuntimeFabricService{
		runtimeFabricClient,
	}
}

type DefaultOrganizationRuntimeFabricService struct {
	runtimeFabricClient organizationRuntimeFabrics.OrganizationRuntimeFabricClient
}

func (this DefaultOrganizationRuntimeFabricService) GetFabrics(orgId, envId, token string) (*[]response.OrganizationFabricResponse, error) {

	resp, err := this.runtimeFabricClient.GetFabrics(orgId, envId, token)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (this DefaultOrganizationRuntimeFabricService) GetFabric(orgId, envId, token, targetId string) (*response.OrganizationFabricResponse, error) {

	resp, err := this.runtimeFabricClient.GetFabric(orgId, envId, token, targetId)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (this DefaultOrganizationRuntimeFabricService) GetTargets(orgId, envId, token string) (*[]response.FabricTargetResponse, error) {

	resp, err := this.runtimeFabricClient.GetTargets(orgId, envId, token)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (this DefaultOrganizationRuntimeFabricService) GetTarget(orgId, envId, token, targetId string) (*response.FabricTargetResponse, error) {

	resp, err := this.runtimeFabricClient.GetTarget(orgId, envId, token, targetId)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
