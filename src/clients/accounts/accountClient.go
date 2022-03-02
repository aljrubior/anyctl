package accounts

import "github.com/aljrubior/anyctl/clients/accounts/response"

type AccountClient interface {
	Login(body []byte) (*response.LoginResponse, error)
	GetProfile(token string) (*response.Profile, error)
	GetOrganization(token, orgId string) (*response.OrganizationResponse, error)
}
