package handlers

import (
	"fmt"
	"github.com/aljrubior/anyctl/errors"
	"github.com/aljrubior/anyctl/managers"
	"github.com/aljrubior/anyctl/model"
	"github.com/aljrubior/anyctl/printers"
)

func NewDefaultConfigHandler(accountManager managers.AccountManager, configManager managers.ConfigManager) DefaultConfigHandler {
	return DefaultConfigHandler{
		configManager:  configManager,
		accountManager: accountManager,
	}
}

type DefaultConfigHandler struct {
	ConfigHandler
	accountManager managers.AccountManager
	configManager  managers.ConfigManager
}

func (this DefaultConfigHandler) Initialize(username, password string) error {

	result, err := this.accountManager.Login(username, password)

	if err != nil {
		return err
	}

	_, err = this.configManager.CreateConfig(username, password, result.Token, result.OrganizationProfile)

	return err
}

func (this DefaultConfigHandler) RefreshAccessToken() error {

	credentials, err := this.configManager.GetCredentials()

	if err != nil {
		return err
	}

	token, err := this.accountManager.RefreshAccessToken(credentials.Username, credentials.Password)

	if err != nil {
		return err
	}

	return this.configManager.UpdateAccessToken(token)
}

func (this DefaultConfigHandler) SetCurrentEnvironment(name string) error {

	envId, options, err := this.configManager.UpdateCurrentEnvironment(name)

	if err != nil {
		return err
	}

	if envId == nil {
		return this.ThrowAnypointEnvironmentNotFoundError(name, options)
	}

	return nil
}

func (this DefaultConfigHandler) PrintCurrentContext() error {

	environmentName, err := this.configManager.GetCurrentEnvironmentName()

	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("Current environment: %s", environmentName))

	return nil
}

func (this DefaultConfigHandler) GetEnvironments() error {

	environments := this.configManager.GetEnvironments()

	printers.NewEnvironmentsPrinter(environments).Print()

	return nil
}

func (this DefaultConfigHandler) ThrowAnypointEnvironmentNotFoundError(environmentName string, options *[]model.Environment) error {
	return errors.NewAnypointEnvironmentNotFoundError(environmentName, options)
}
