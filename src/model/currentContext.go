package model

type currentContext struct {
	EnvironmentId      string `yaml:"environmentId"`
	AuthorizationToken string `yaml:"authorizationToken"`
}
