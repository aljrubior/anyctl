package managers

import (
	"github.com/aljrubior/anyctl/managers/entities"
	"github.com/aljrubior/anyctl/services"
)

func NewDefaultLoginManager(accountService services.AccountService) *DefaultLoginManager {
	return &DefaultLoginManager{
		accountService: accountService,
	}
}

type DefaultLoginManager struct {
	LoginManager
	accountService services.AccountService
}

func (this DefaultLoginManager) Login(username, password string) (*entities.LoginEntity, error) {

	token, err := this.accountService.GetAuthorizationToken(username, password)

	if err != nil {
		return nil, err
	}

	profile, err := this.accountService.GetProfile(token)

	if err != nil {
		return nil, err
	}

	return entities.NewLoginEntityBuilder(token, profile).Build(), nil
}

func (this DefaultLoginManager) RefreshAccessToken(username, password string) (string, error) {

	token, err := this.accountService.GetAuthorizationToken(username, password)

	if err != nil {
		return "", err
	}

	return token, nil
}
