package services

import (
	response2 "github.com/aljrubior/anyctl/clients/fabrics/response"
	"github.com/aljrubior/anyctl/clients/targets"
	"github.com/aljrubior/anyctl/clients/targets/response"
)

func NewDefaultTargetService(targetClient targets.TargetClient) *DefaultTargetService {
	return &DefaultTargetService{
		targetClient: targetClient,
	}
}

type DefaultTargetService struct {
	targetClient targets.TargetClient
}

func (this *DefaultTargetService) GetTargets(orgId, envId, token string) (*[]response.TargetResponse, error) {

	resp, err := this.targetClient.GetTargets(orgId, envId, token)

	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

func (this DefaultTargetService) GetFabrics(name, token string) (*[]response2.FabricResponse, error) {

	resp, err := this.targetClient.GetFabrics(name, token)

	if err != nil {
		return nil, err
	}

	return &resp.Content, nil
}
