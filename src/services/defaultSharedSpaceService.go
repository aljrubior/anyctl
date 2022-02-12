package services

import (
	"github.com/aljrubior/anyctl/clients/sharedspaces"
	"github.com/aljrubior/anyctl/clients/sharedspaces/response"
)

func NewDefaultSharedSpaceService(sharedSpaceClient sharedspaces.SharedSpaceClient) *DefaultSharedSpaceService {
	return &DefaultSharedSpaceService{
		sharedSpaceClient,
	}
}

type DefaultSharedSpaceService struct {
	sharedSpaceClient sharedspaces.SharedSpaceClient
}

func (this DefaultSharedSpaceService) GetSharedSpaces(token string) (*[]response.SharedSpaceResponse, error) {

	return this.sharedSpaceClient.GetSharedSpaces(token)
}

func (this DefaultSharedSpaceService) GetSharedSpace(token, targetId string) (*response.SharedSpaceResponse, error) {

	return this.sharedSpaceClient.GetSharedSpace(token, targetId)
}
