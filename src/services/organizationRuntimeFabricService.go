package services

import (
	"github.com/aljrubior/anyctl/clients/organizationRuntimeFabrics/response"
)

type OrganizationRuntimeFabricService interface {
	GetFabrics(orgId, envId, token string) (*[]response.OrganizationFabricResponse, error)
	GetFabric(orgId, envId, token, targetId string) (*response.OrganizationFabricResponse, error)
	GetTargets(orgId, envId, token string) (*[]response.FabricTargetResponse, error)
	GetTarget(orgId, envId, token, targetId string) (*response.FabricTargetResponse, error)
}
