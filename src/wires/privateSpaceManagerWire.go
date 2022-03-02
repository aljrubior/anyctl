//go:build wireinject
// +build wireinject

package wires

import (
	"github.com/aljrubior/anyctl/clients/privateSpaces"
	"github.com/aljrubior/anyctl/conf"
	"github.com/aljrubior/anyctl/managers"
	"github.com/aljrubior/anyctl/services"
	"github.com/google/wire"
)

func InitializePrivateSpaceManager(config conf.PrivateSpaceClientConfig) (managers.PrivateSpaceManager, error) {

	wire.Build(
		privateSpaces.NewDefaultPrivateSpaceClient,
		services.NewDefaultPrivateSpaceService,
		managers.NewDefaultPrivateSpaceManager,
		wire.Bind(new(privateSpaces.PrivateSpaceClient), new(privateSpaces.DefaultPrivateSpaceClient)),
		wire.Bind(new(services.PrivateSpaceService), new(services.DefaultPrivateSpaceService)),
		wire.Bind(new(managers.PrivateSpaceManager), new(managers.DefaultPrivateSpaceManager)),
	)

	return managers.DefaultPrivateSpaceManager{}, nil
}
