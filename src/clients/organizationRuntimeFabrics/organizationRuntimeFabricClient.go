package organizationRuntimeFabrics

import (
	"github.com/aljrubior/anyctl/clients/organizationRuntimeFabrics/response"
)

type OrganizationRuntimeFabricClient interface {
	GetFabrics(orgId, envId, token string) (*[]response.OrganizationFabricResponse, error)
	GetFabric(orgId, envId, token, targetId string) (*response.OrganizationFabricResponse, error)
	GetTarget(orgId, envId, token, targetId string) (*response.FabricTargetResponse, error)
	GetTargets(orgId, envId, token string) (*[]response.FabricTargetResponse, error)
}
