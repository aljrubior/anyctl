package organizationPrivateSpaces

import (
	"github.com/aljrubior/anyctl/clients/organizationPrivateSpaces/response"
)

type PrivateSpaceClient interface {
	GetPrivateSpaces(orgId, envId, token string) (*response.OrganizationPrivateSpacesResponse, error)
	GetPrivateSpace(orgId, envId, token, targetId string) (*response.OrganizationPrivateSpaceResponse, error)
	GetFabrics(orgId, envId, token, privateSpaceId string) (*[]response.OrganizationPrivateSpaceFabricResponse, error)
}
