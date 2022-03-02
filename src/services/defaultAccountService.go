package services

import (
	"encoding/json"
	"github.com/aljrubior/anyctl/clients/accounts"
	"github.com/aljrubior/anyctl/clients/accounts/response"
	"github.com/aljrubior/anyctl/managers/requests"
)

func NewDefaultAccountService(accountClient accounts.AccountClient) DefaultAccountService {
	return DefaultAccountService{
		accountClient,
	}
}

type DefaultAccountService struct {
	accountClient accounts.AccountClient
}

func (this DefaultAccountService) GetOrganization(token, orgId string) (*response.OrganizationResponse, error) {

	return this.accountClient.GetOrganization(token, orgId)
}

func (this DefaultAccountService) Login(request requests.LoginRequest) (*string, error) {

	body, err := json.Marshal(request)

	if err != nil {
		return nil, err
	}

	resp, err := this.accountClient.Login(body)

	if err != nil {
		return nil, err
	}

	return &resp.AccessToken, nil
}

func (this DefaultAccountService) GetProfile(accessToken string) (*response.Profile, error) {

	return this.accountClient.GetProfile(accessToken)
}
