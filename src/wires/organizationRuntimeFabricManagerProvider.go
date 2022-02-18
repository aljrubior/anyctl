//go:build wireinject
// +build wireinject

package wires

import (
	"github.com/aljrubior/anyctl/clients/organizationRuntimeFabrics"
	"github.com/aljrubior/anyctl/conf"
	"github.com/aljrubior/anyctl/managers"
	"github.com/aljrubior/anyctl/services"
	"github.com/google/wire"
)

func InitializeOrganizationRuntimeFabricManager(config conf.RuntimeFabricClientConfig) (managers.OrganizationRuntimeFabricManager, error) {

	wire.Build(
		organizationRuntimeFabrics.NewDefaultOrganizationRuntimeFabricClient,
		services.NewDefaultOrganizationRuntimeFabricService,
		managers.NewDefaultOrganizationRuntimeFabricManager,
		wire.Bind(new(organizationRuntimeFabrics.OrganizationRuntimeFabricClient), new(organizationRuntimeFabrics.DefaultOrganizationRuntimeFabricClient)),
		wire.Bind(new(services.OrganizationRuntimeFabricService), new(services.DefaultOrganizationRuntimeFabricService)),
		wire.Bind(new(managers.OrganizationRuntimeFabricManager), new(managers.DefaultOrganizationRuntimeFabricManager)),
	)

	return managers.DefaultOrganizationRuntimeFabricManager{}, nil
}
