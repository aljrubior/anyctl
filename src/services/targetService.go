package services

import (
	response2 "github.com/aljrubior/anyctl/clients/fabrics/response"
	"github.com/aljrubior/anyctl/clients/targets/response"
)

type TargetService interface {
	GetTargets(orgId, envId, token string) (*[]response.TargetResponse, error)
	GetFabrics(name, token string) (*[]response2.FabricResponse, error)
}
