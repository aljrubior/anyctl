//go:build wireinject
// +build wireinject

package wires

import (
	"github.com/aljrubior/anyctl/clients/targets"
	"github.com/aljrubior/anyctl/conf"
	"github.com/aljrubior/anyctl/managers"
	"github.com/aljrubior/anyctl/services"
	"github.com/google/wire"
)

func InitializeTargetManager(config conf.TargetClientConfig) (managers.TargetManager, error) {

	wire.Build(
		targets.NewDefaultTargetClient,
		services.NewDefaultTargetService,
		managers.NewDefaultTargetManager,
		wire.Bind(new(targets.TargetClient), new(targets.DefaultTargetClient)),
		wire.Bind(new(services.TargetService), new(services.DefaultTargetService)),
		wire.Bind(new(managers.TargetManager), new(managers.DefaultTargetManager)),
	)

	return managers.DefaultTargetManager{}, nil
}
