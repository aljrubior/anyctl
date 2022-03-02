package targets

import (
	"github.com/aljrubior/anyctl/clients/targets/response"
)

type TargetClient interface {
	GetTargets(orgId, envId, token string) (*response.TargetsResponse, error)
}
