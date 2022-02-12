package accounts

import "github.com/aljrubior/anyctl/clients/accounts/response"

type AccountClient interface {
	GetAuthorizationToken(username, password string) (*response.LoginResponse, error)
	GetProfile(token string) (*response.Profile, error)
	GetOrganization(token, orgId string) (*response.OrganizationResponse, error)
}
