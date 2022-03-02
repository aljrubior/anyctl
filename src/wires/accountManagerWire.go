//go:build wireinject
// +build wireinject

package wires

import (
	"github.com/aljrubior/anyctl/clients/accounts"
	"github.com/aljrubior/anyctl/conf"
	"github.com/aljrubior/anyctl/managers"
	"github.com/aljrubior/anyctl/services"
	"github.com/google/wire"
)

func InitializeAccountManager(config conf.AccountClientConfig) (managers.AccountManager, error) {

	wire.Build(
		accounts.NewDefaultAccountClient,
		services.NewDefaultAccountService,
		managers.NewDefaultAccountManager,
		wire.Bind(new(accounts.AccountClient), new(accounts.DefaultAccountClient)),
		wire.Bind(new(services.AccountService), new(services.DefaultAccountService)),
		wire.Bind(new(managers.AccountManager), new(managers.DefaultAccountManager)),
	)

	return managers.DefaultAccountManager{}, nil
}
