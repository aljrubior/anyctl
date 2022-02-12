package services

import (
	response2 "github.com/aljrubior/anyctl/clients/fabrics/response"
	"github.com/aljrubior/anyctl/clients/privateSpaces"
	"github.com/aljrubior/anyctl/clients/privateSpaces/response"
)

func NewDefaultPrivateSpaceService(privateSpaceClient privateSpaces.PrivateSpaceClient) *DefaultPrivateSpaceService {
	return &DefaultPrivateSpaceService{
		privateSpaceClient,
	}
}

type DefaultPrivateSpaceService struct {
	privateSpaceClient privateSpaces.PrivateSpaceClient
}

func (this DefaultPrivateSpaceService) GetPrivateSpaces(token string) (*[]response.PrivateSpaceResponse, error) {

	resp, err := this.privateSpaceClient.GetPrivateSpaces(token)

	if err != nil {
		return nil, err
	}

	return &resp.Content, nil
}

func (this DefaultPrivateSpaceService) GetPrivateSpace(token, targetId string) (*response.PrivateSpaceResponse, error) {

	return this.privateSpaceClient.GetPrivateSpace(token, targetId)
}

func (this DefaultPrivateSpaceService) GetPrivateSpacesByNameOrId(token, privateSpaceId string) (*[]response.PrivateSpaceResponse, error) {

	resp, err := this.privateSpaceClient.GetPrivateSpacesByNameOrId(token, privateSpaceId)

	if err != nil {
		return nil, err
	}

	return &resp.Content, nil
}

func (this DefaultPrivateSpaceService) GetFabrics(token, privateSpaceId string) (*[]response2.FabricResponse, error) {

	resp, err := this.privateSpaceClient.GetFabrics(token, privateSpaceId)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (this DefaultPrivateSpaceService) GetFabric(token, privateSpaceId, fabricId string) (*response2.FabricResponse, error) {

	return this.privateSpaceClient.GetFabric(token, privateSpaceId, fabricId)
}
