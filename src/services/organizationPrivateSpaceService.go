package services

import (
	"github.com/aljrubior/anyctl/clients/organizationPrivateSpaces/response"
)

type OrganizationPrivateSpaceService interface {
	GetPrivateSpaces(orgId, envId, token string) (*[]response.OrganizationPrivateSpaceResponse, error)
	GetPrivateSpace(orgId, envId, token, targetId string) (*response.OrganizationPrivateSpaceResponse, error)
	GetFabrics(orgId, envId, token, privateSpaceId string) (*[]response.OrganizationPrivateSpaceFabricResponse, error)
}
