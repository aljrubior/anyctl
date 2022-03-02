package services

import (
	"github.com/aljrubior/anyctl/clients/targets/response"
)

type TargetService interface {
	GetTargets(orgId, envId, token string) (*[]response.TargetResponse, error)
}
