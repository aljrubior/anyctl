package handlers

import (
	"fmt"
	"github.com/aljrubior/anyctl/managers"
)

func NewDefaultLoginHandler(loginManager managers.AccountManager, configManager managers.ConfigManager) *DefaultLoginHandler {
	return &DefaultLoginHandler{
		loginManager:  loginManager,
		configManager: configManager,
	}
}

type DefaultLoginHandler struct {
	LoginHandler
	loginManager  managers.AccountManager
	configManager managers.ConfigManager
}

func (this DefaultLoginHandler) Login(username, password string) error {

	loginEntity, err := this.loginManager.Login(username, password)

	if err != nil {
		return err
	}

	this.configManager.CreateConfig(username, password, loginEntity.Token, loginEntity.OrganizationProfile)

	anyconfigPath, err := this.configManager.GetAnyConfigFilePath()

	if err != nil {
		return err
	}

	println(fmt.Sprintf("Configuration file '%s' created.", anyconfigPath))

	return err
}
