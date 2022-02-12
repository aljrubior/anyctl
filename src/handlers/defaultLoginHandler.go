package handlers

import (
	"github.com/aljrubior/anyctl/managers"
)

func NewDefaultLoginHandler(loginManager managers.LoginManager, configManager managers.ConfigManager) *DefaultLoginHandler {
	return &DefaultLoginHandler{
		loginManager:  loginManager,
		configManager: configManager,
	}
}

type DefaultLoginHandler struct {
	LoginHandler
	loginManager  managers.LoginManager
	configManager managers.ConfigManager
}

func (this DefaultLoginHandler) Login(username, password string) error {

	loginEntity, err := this.loginManager.Login(username, password)

	if err != nil {
		return err
	}

	this.configManager.CreateConfig(username, password, loginEntity.Token, loginEntity.OrganizationProfile)

	return err
}
