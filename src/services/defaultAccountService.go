package services

import (
	"github.com/aljrubior/anyctl/clients/accounts"
	"github.com/aljrubior/anyctl/clients/accounts/response"
)

func NewDefaultAccountService(accountClient accounts.AccountClient) *DefaultAccountService {
	return &DefaultAccountService{
		accountClient,
	}
}

type DefaultAccountService struct {
	accountClient accounts.AccountClient
}

func (this *DefaultAccountService) GetOrganization(token, orgId string) (*response.OrganizationResponse, error) {

	return this.accountClient.GetOrganization(token, orgId)
}

func (this *DefaultAccountService) GetAuthorizationToken(username, password string) (string, error) {

	resp, err := this.accountClient.GetAuthorizationToken(username, password)

	if err != nil {
		return "", err
	}

	return resp.AccessToken, nil
}

func (this *DefaultAccountService) GetProfile(accessToken string) (*response.Profile, error) {

	return this.accountClient.GetProfile(accessToken)
}
