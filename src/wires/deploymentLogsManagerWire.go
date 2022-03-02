//go:build wireinject
// +build wireinject

package wires

import (
	"github.com/aljrubior/anyctl/clients/deploymentLogs"
	"github.com/aljrubior/anyctl/conf"
	"github.com/aljrubior/anyctl/managers"
	"github.com/aljrubior/anyctl/services"
	"github.com/google/wire"
)

func InitializeDeploymentLogsManager(config conf.DeploymentLogsClientConfig) (managers.DeploymentLogsManager, error) {

	wire.Build(
		deploymentLogs.NewDefaultDeploymentLogsClient,
		services.NewDefaultDeploymentLogsService,
		managers.NewDefaultDeploymentLogsManager,
		wire.Bind(new(deploymentLogs.DeploymentLogsClient), new(deploymentLogs.DefaultDeploymentLogsClient)),
		wire.Bind(new(services.DeploymentLogsService), new(services.DefaultDeploymentLogsService)),
		wire.Bind(new(managers.DeploymentLogsManager), new(managers.DefaultDeploymentLogsManager)),
	)

	return managers.DefaultDeploymentLogsManager{}, nil
}
