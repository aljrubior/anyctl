//go:build wireinject
// +build wireinject

package wires

import (
	"github.com/aljrubior/anyctl/clients/fabrics"
	"github.com/aljrubior/anyctl/conf"
	"github.com/aljrubior/anyctl/managers"
	"github.com/aljrubior/anyctl/services"
	"github.com/google/wire"
)

func InitializeFabricManager(config conf.FabricClientConfig) (managers.FabricManager, error) {

	wire.Build(
		fabrics.NewDefaultFabricClient,
		services.NewDefaultFabricService,
		managers.NewDefaultFabricManager,
		wire.Bind(new(fabrics.FabricClient), new(fabrics.DefaultFabricClient)),
		wire.Bind(new(services.FabricService), new(services.DefaultFabricService)),
		wire.Bind(new(managers.FabricManager), new(managers.DefaultFabricManager)),
	)

	return managers.DefaultFabricManager{}, nil
}
