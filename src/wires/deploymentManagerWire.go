//go:build wireinject
// +build wireinject

package wires

import (
	"github.com/aljrubior/anyctl/clients/deployments"
	"github.com/aljrubior/anyctl/conf"
	"github.com/aljrubior/anyctl/managers"
	"github.com/aljrubior/anyctl/services"
	"github.com/google/wire"
)

func InitializeDeploymentManager(config conf.DeploymentClientConfig, assetManager managers.AssetManager) (managers.DeploymentManager, error) {

	wire.Build(
		deployments.NewDefaultDeploymentClient,
		services.NewDefaultDeploymentService,
		managers.NewDefaultDeploymentManager,
		wire.Bind(new(deployments.DeploymentClient), new(deployments.DefaultDeploymentClient)),
		wire.Bind(new(services.DeploymentService), new(services.DefaultDeploymentService)),
		wire.Bind(new(managers.DeploymentManager), new(managers.DefaultDeploymentManager)),
	)

	return managers.DefaultDeploymentManager{}, nil
}
