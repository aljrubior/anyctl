//go:build wireinject
// +build wireinject

package wires

import (
	"github.com/aljrubior/anyctl/conf"
	"github.com/aljrubior/anyctl/managers"
	"github.com/google/wire"
)

func InitializeConfigManager() managers.ConfigManager {
	wire.Build(
		conf.NewAppConfig,
		managers.NewDefaultConfigManager,
		wire.Bind(new(managers.ConfigManager), new(managers.DefaultConfigManager)),
	)

	return managers.DefaultConfigManager{}
}
