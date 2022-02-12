package managers

import (
	"github.com/aljrubior/anyctl/clients/accounts/response"
	"github.com/aljrubior/anyctl/managers/entities"
	"github.com/aljrubior/anyctl/model"
)

type ConfigManager interface {
	GetAnyConfigFilePath() (string, error)
	CreateConfig(username, password, token string, profile response.Profile) (*model.Config, error)
	UpdateAccessToken(token string) error
	UpdateCurrentEnvironment(name string) (*string, *[]model.Environment, error)
	GetCurrentEnvironmentName() (string, error)
	GetEnvironments() *[]model.Environment
	GetOrganizationId() (string, error)
	GetCurrentEnvironmentId() (string, error)
	GetAuthorizationToken() (string, error)
	GetCredentials() (model.Credentials, error)
	GetCurrentContext() (*entities.CurrentContextEntity, error)
}
