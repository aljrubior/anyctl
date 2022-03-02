//go:build wireinject
// +build wireinject

package wires

import (
	"github.com/aljrubior/anyctl/clients/assets"
	"github.com/aljrubior/anyctl/conf"
	"github.com/aljrubior/anyctl/managers"
	"github.com/aljrubior/anyctl/services"
	"github.com/google/wire"
)

func InitializeAssetManager(config conf.AssetClientConfig) (managers.AssetManager, error) {

	wire.Build(
		assets.NewDefaultAssetClient,
		services.NewDefaultAssetService,
		managers.NewDefaultAssetManager,
		wire.Bind(new(assets.AssetClient), new(assets.DefaultAssetClient)),
		wire.Bind(new(services.AssetService), new(services.DefaultAssetService)),
		wire.Bind(new(managers.AssetManager), new(managers.DefaultAssetManager)),
	)

	return managers.DefaultAssetManager{}, nil
}
