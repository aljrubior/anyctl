package targets

import (
	response2 "github.com/aljrubior/anyctl/clients/fabrics/response"
	"github.com/aljrubior/anyctl/clients/targets/response"
)

type TargetClient interface {
	GetTargets(orgId, envId, token string) (*response.TargetsResponse, error)

	GetFabrics(name, token string) (*response2.FabricsResponse, error)
}
