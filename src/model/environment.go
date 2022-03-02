package model

type Environment struct {
	Id             string `yaml:"id"`
	Name           string `yaml:"name"`
	Kind           string `yaml:"kind"`
	OrganizationId string `yaml:"organizationId"`
}
