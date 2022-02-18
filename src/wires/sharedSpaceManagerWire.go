//go:build wireinject
// +build wireinject

package wires

import (
	"github.com/aljrubior/anyctl/clients/sharedspaces"
	"github.com/aljrubior/anyctl/conf"
	"github.com/aljrubior/anyctl/managers"
	"github.com/aljrubior/anyctl/services"
	"github.com/google/wire"
)

func InitializeSharedSpaceManager(config conf.SharedSpaceClientConfig) (managers.SharedSpaceManager, error) {

	wire.Build(
		sharedspaces.NewDefaultSharedSpaceClient,
		services.NewDefaultSharedSpaceService,
		managers.NewDefaultSharedSpaceManager,
		wire.Bind(new(sharedspaces.SharedSpaceClient), new(sharedspaces.DefaultSharedSpaceClient)),
		wire.Bind(new(services.SharedSpaceService), new(services.DefaultSharedSpaceService)),
		wire.Bind(new(managers.SharedSpaceManager), new(managers.DefaultSharedSpaceManager)),
	)

	return managers.DefaultSharedSpaceManager{}, nil
}
