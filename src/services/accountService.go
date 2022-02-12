package services

import "github.com/aljrubior/anyctl/clients/accounts/response"

type AccountService interface {
	GetAuthorizationToken(username, password string) (string, error)
	GetProfile(accessToken string) (*response.Profile, error)
	GetOrganization(token, orgId string) (*response.OrganizationResponse, error)
}
