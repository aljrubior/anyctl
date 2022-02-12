package managers

import (
	"fmt"
	"github.com/aljrubior/anyctl/clients/accounts/response"
	"github.com/aljrubior/anyctl/conf"
	"github.com/aljrubior/anyctl/managers/entities"
	"github.com/aljrubior/anyctl/model"
	"github.com/aljrubior/anyctl/utils"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func NewDefaultConfigManager(appConfig conf.AppConfig) *DefaultConfigManager {

	return &DefaultConfigManager{
		configDir:  appConfig.ConfigDir(),
		configFile: appConfig.AnypointConfigFile(),
	}
}

type DefaultConfigManager struct {
	ConfigManager
	region     model.Region
	config     model.Config
	configFile string
	configDir  string
}

func (this DefaultConfigManager) GetAnyConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s/%s", homeDir, this.configDir, this.configFile), nil
}

func (this DefaultConfigManager) CreateConfig(username, password, token string, profile response.Profile) (*model.Config, error) {
	config := model.NewConfig()

	this.appendEnvironments(config, profile)
	defaultEnvironment := profile.GetDefaultEnvironment()
	config.SetCurrentContext(defaultEnvironment.Id, token)
	config.SetCredentials(username, password)
	config.SetOrgnizationId(profile.Organization.Id)

	data, err := yaml.Marshal(&config)
	if err != nil {
		return nil, err
	}

	anypointDir, err := this.getConfigDir()
	if err != nil {
		return nil, err
	}

	this.createConfigDir(anypointDir)

	configFilePath := fmt.Sprintf("%s/%s", anypointDir, this.configFile)

	err = utils.WriteFile(configFilePath, data)

	return config, nil
}

func (this DefaultConfigManager) UpdateAccessToken(token string) error {
	err := this.loadConfig()

	if err != nil {
		return err
	}

	this.config.CurrentContext.AuthorizationToken = token

	return this.saveConfig()
}

func (this DefaultConfigManager) UpdateCurrentEnvironment(name string) (*string, *[]model.Environment, error) {

	if err := this.loadConfig(); err != nil {
		return nil, nil, err
	}

	envId := this.getEnvironmentId(name)

	if len(envId) == 0 {
		return nil, this.GetEnvironments(), nil
	}

	this.config.SetCurrentEnvironment(envId)

	if err := this.saveConfig(); err != nil {
		return nil, nil, err
	}

	return &envId, nil, nil
}

func (this DefaultConfigManager) saveConfig() error {

	filePath, err := this.GetAnyConfigFilePath()

	if err != nil {
		return err
	}

	data, err := yaml.Marshal(&this.config)

	if err != nil {
		return err
	}

	if err = utils.WriteFile(filePath, data); err != nil {
		return err
	}

	return nil
}

func (this DefaultConfigManager) appendEnvironments(config *model.Config, profile response.Profile) {
	for _, v := range profile.Organization.Environments {
		config.AddEnvironment(v.Id, v.Name, v.Kind, v.OrganizationId, this.region.Name)
	}
}

func (this DefaultConfigManager) createConfigDir(configDir string) error {
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		err := os.Mkdir(configDir, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}

func (this DefaultConfigManager) getConfigDir() (string, error) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s", homeDir, this.configDir), nil
}

func (this *DefaultConfigManager) loadConfig() error {

	anypointHome, err := this.getConfigDir()

	if err != nil {
		return err
	}

	filePath := fmt.Sprintf("%s/%s", anypointHome, this.configFile)

	content, err := ioutil.ReadFile(filePath)

	if err != nil {
		return err
	}

	config := model.Config{}

	err = yaml.Unmarshal(content, &config)

	if err != nil {
		return err
	}

	this.config = config

	return nil
}

func (this DefaultConfigManager) getEnvironmentId(fromName string) string {

	for _, v := range this.config.Environments {
		if v.Name == fromName {
			return v.Id
		}
	}

	return ""
}

func (this DefaultConfigManager) GetEnvironments() *[]model.Environment {

	return &this.config.Environments
}

func (this DefaultConfigManager) getEnvironmentName(fromId string) string {

	for _, v := range this.config.Environments {
		if v.Id == fromId {
			return v.Name
		}
	}

	return ""
}

func (this DefaultConfigManager) GetCurrentEnvironmentName() (string, error) {
	if err := this.loadConfig(); err != nil {
		return "", err
	}
	return this.getEnvironmentName(this.config.CurrentContext.EnvironmentId), nil
}

func (this DefaultConfigManager) GetOrganizationId() (string, error) {

	if err := this.loadConfig(); err != nil {
		return "", err
	}
	return this.config.OrganizationId, nil
}

func (this DefaultConfigManager) GetCurrentEnvironmentId() (string, error) {

	if err := this.loadConfig(); err != nil {
		return "", err
	}
	return this.config.CurrentContext.EnvironmentId, nil
}

func (this DefaultConfigManager) GetAuthorizationToken() (string, error) {

	if err := this.loadConfig(); err != nil {
		return "", err
	}

	return this.config.CurrentContext.AuthorizationToken, nil
}

func (this DefaultConfigManager) GetCredentials() (model.Credentials, error) {

	if err := this.loadConfig(); err != nil {
		return model.Credentials{}, err
	}

	return this.config.Credentials, nil
}

func (this DefaultConfigManager) GetCurrentContext() (*entities.CurrentContextEntity, error) {

	if err := this.loadConfig(); err != nil {
		return nil, err
	}

	return &entities.CurrentContextEntity{
		OrganizationId:     this.config.OrganizationId,
		EnvironmentId:      this.config.CurrentContext.EnvironmentId,
		AuthorizationToken: this.config.CurrentContext.AuthorizationToken,
	}, nil
}
