//go:build wireinject
// +build wireinject

package wires

import (
	"github.com/aljrubior/anyctl/clients/organizationPrivateSpaces"
	"github.com/aljrubior/anyctl/conf"
	"github.com/aljrubior/anyctl/managers"
	"github.com/aljrubior/anyctl/services"
	"github.com/google/wire"
)

func InitializeOrganizationPrivateSpaceManager(config conf.RuntimeFabricClientConfig) (managers.OrganizationPrivateSpaceManager, error) {

	wire.Build(
		organizationPrivateSpaces.NewOrganizationDefaultPrivateSpaceClient,
		services.NewDefaultOrganizationPrivateSpaceService,
		managers.NewDefaultOrganizationPrivateSpaceManager,
		wire.Bind(new(organizationPrivateSpaces.OrganizationPrivateSpaceClient), new(organizationPrivateSpaces.DefaultOrganizationPrivateSpaceClient)),
		wire.Bind(new(services.OrganizationPrivateSpaceService), new(services.DefaultOrganizationPrivateSpaceService)),
		wire.Bind(new(managers.OrganizationPrivateSpaceManager), new(managers.DefaultOrganizationPrivateSpaceManager)),
	)

	return managers.DefaultOrganizationPrivateSpaceManager{}, nil
}
