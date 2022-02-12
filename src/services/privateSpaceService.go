package services

import (
	response2 "github.com/aljrubior/anyctl/clients/fabrics/response"
	"github.com/aljrubior/anyctl/clients/privateSpaces/response"
)

type PrivateSpaceService interface {
	GetPrivateSpaces(token string) (*[]response.PrivateSpaceResponse, error)
	GetPrivateSpace(token, privateSpaceId string) (*response.PrivateSpaceResponse, error)
	GetPrivateSpacesByNameOrId(token string, privateSpaceId string) (*[]response.PrivateSpaceResponse, error)

	GetFabrics(token, privateSpaceId string) (*[]response2.FabricResponse, error)
	GetFabric(token, privateSpaceId, fabricId string) (*response2.FabricResponse, error)
}
