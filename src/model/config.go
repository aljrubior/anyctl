package model

func NewConfig() *Config {
	return &Config{
		ApiVersion:     "v1",
		Kind:           "Config",
		OrganizationId: "",
		Credentials:    Credentials{},
		Environments:   []Environment{},
		CurrentContext: currentContext{},
	}
}

type Config struct {
	ApiVersion     string         `yaml:"apiVersion"`
	Kind           string         `yaml:"kind"`
	Credentials    Credentials    `yaml:"credentials"`
	OrganizationId string         `yaml:"organizationId"`
	Environments   []Environment  `yaml:"environments"`
	CurrentContext currentContext `yaml:"currentContext"`
}

func (c Config) GetApiVersion() string {
	return c.ApiVersion
}

func (c Config) GetEnvironments() []Environment {
	return c.Environments
}

func (c *Config) AddEnvironment(id, name, kind, organizationId, region string) {
	c.Environments = append(c.Environments, Environment{
		id,
		name,
		kind,
		organizationId,
	})
}

func (c *Config) SetCurrentContext(environmentId, authorizationToken string) {
	c.CurrentContext = currentContext{environmentId, authorizationToken}
}

func (c Config) GetCurrentContext() currentContext {
	return c.CurrentContext
}

func (c *Config) SetCredentials(username, password string) {
	c.Credentials = Credentials{username, password}
}

func (c *Config) SetOrgnizationId(organizationId string) {
	c.OrganizationId = organizationId
}

func (c *Config) SetCurrentEnvironment(id string) {
	c.CurrentContext.EnvironmentId = id
}

func (c *Config) SetAccessToken(token string) {
	c.CurrentContext.AuthorizationToken = token
}
