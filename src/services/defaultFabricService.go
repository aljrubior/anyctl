package services

import (
	"github.com/aljrubior/anyctl/clients/fabrics"
	"github.com/aljrubior/anyctl/clients/fabrics/response"
)

func NewDefaultFabricService(fabricClient fabrics.FabricClient) *DefaultFabricService {
	return &DefaultFabricService{
		fabricClient,
	}
}

type DefaultFabricService struct {
	fabricClient fabrics.FabricClient
}

func (this DefaultFabricService) GetFabrics(token string) (*[]response.FabricResponse, error) {

	return this.fabricClient.GetFabrics(token)
}

func (this DefaultFabricService) GetFabric(token, fabricId string) (*response.FabricResponse, error) {

	return this.fabricClient.GetFabric(token, fabricId)
}

func (this DefaultFabricService) GetFabricsByNameOrId(token, fabricId string) (*[]response.FabricResponse, error) {

	resp, err := this.fabricClient.GetFabricsByNameOrId(token, fabricId)

	if err != nil {
		return nil, err
	}

	return &resp.Content, nil
}
