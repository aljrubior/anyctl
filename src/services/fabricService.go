package services

import (
	"github.com/aljrubior/anyctl/clients/fabrics/response"
)

type FabricService interface {
	GetFabrics(token string) (*[]response.FabricResponse, error)
	GetFabric(token, fabricId string) (*response.FabricResponse, error)
	GetFabricsByNameOrId(token, fabricId string) (*[]response.FabricResponse, error)
}
