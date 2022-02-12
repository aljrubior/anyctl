package fabrics

import (
	"github.com/aljrubior/anyctl/clients/fabrics/response"
)

type FabricClient interface {
	GetFabrics(token string) (*[]response.FabricResponse, error)
	GetFabric(token, fabricId string) (*response.FabricResponse, error)
	GetFabricsByNameOrId(token string, fabricId string) (*response.FabricsResponse, error)
}
