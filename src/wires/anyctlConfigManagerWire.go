//go:build wireinject
// +build wireinject

package wires

import (
	"github.com/aljrubior/anyctl/conf"
	"github.com/aljrubior/anyctl/managers"
	"github.com/google/wire"
)

func InitializeAnyctlConfigManager() (managers.AnyctlConfigManager, error) {
	wire.Build(
		conf.NewAppConfig,
		managers.NewAnyctlConfigManager)
	return managers.AnyctlConfigManager{}, nil
}
