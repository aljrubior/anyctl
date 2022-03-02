package services

import (
	"github.com/aljrubior/anyctl/clients/organizationPrivateSpaces"
	"github.com/aljrubior/anyctl/clients/organizationPrivateSpaces/response"
)

func NewDefaultOrganizationPrivateSpaceService(privateSpaceClient organizationPrivateSpaces.OrganizationPrivateSpaceClient) DefaultOrganizationPrivateSpaceService {
	return DefaultOrganizationPrivateSpaceService{
		privateSpaceClient,
	}
}

type DefaultOrganizationPrivateSpaceService struct {
	privateSpaceClient organizationPrivateSpaces.OrganizationPrivateSpaceClient
}

func (this DefaultOrganizationPrivateSpaceService) GetPrivateSpaces(orgId, envId, token string) (*[]response.OrganizationPrivateSpaceResponse, error) {

	resp, err := this.privateSpaceClient.GetPrivateSpaces(orgId, envId, token)

	if err != nil {
		return nil, err
	}

	return &resp.Content, nil
}

func (this DefaultOrganizationPrivateSpaceService) GetFabrics(orgId, envId, token, privateSpaceId string) (*[]response.OrganizationPrivateSpaceFabricResponse, error) {

	resp, err := this.privateSpaceClient.GetFabrics(orgId, envId, token, privateSpaceId)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (this DefaultOrganizationPrivateSpaceService) GetPrivateSpace(orgId, envId, token, targetId string) (*response.OrganizationPrivateSpaceResponse, error) {

	resp, err := this.privateSpaceClient.GetPrivateSpace(orgId, envId, token, targetId)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
