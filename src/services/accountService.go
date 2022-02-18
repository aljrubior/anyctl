package services

import (
	"github.com/aljrubior/anyctl/clients/accounts/response"
	"github.com/aljrubior/anyctl/managers/requests"
)

type AccountService interface {
	Login(request requests.LoginRequest) (*string, error)
	GetProfile(accessToken string) (*response.Profile, error)
	GetOrganization(token, orgId string) (*response.OrganizationResponse, error)
}
