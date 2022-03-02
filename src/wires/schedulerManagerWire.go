//go:build wireinject
// +build wireinject

package wires

import (
	"github.com/aljrubior/anyctl/clients/schedulers"
	"github.com/aljrubior/anyctl/conf"
	"github.com/aljrubior/anyctl/managers"
	"github.com/aljrubior/anyctl/services"
	"github.com/google/wire"
)

func InitializeSchedulerManager(config conf.SchedulerClientConfig, manager managers.DeploymentManager) (managers.SchedulerManager, error) {

	wire.Build(
		schedulers.NewDefaultSchedulerClient,
		services.NewDefaultSchedulerService,
		managers.NewDefaultSchedulerManager,
		wire.Bind(new(schedulers.SchedulerClient), new(schedulers.DefaultSchedulerClient)),
		wire.Bind(new(services.SchedulerService), new(services.DefaultSchedulerService)),
		wire.Bind(new(managers.SchedulerManager), new(managers.DefaultSchedulerManager)),
	)

	return managers.DefaultSchedulerManager{}, nil
}
