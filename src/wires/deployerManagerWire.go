//go:build wireinject
// +build wireinject

package wires

import (
	"github.com/aljrubior/anyctl/managers"
	"github.com/google/wire"
)

func InitializeDeployerManager(deploymentManager managers.DeploymentManager, assetManager managers.AssetManager, fabricManager managers.OrganizationRuntimeFabricManager) (managers.DeployerManager, error) {

	wire.Build(
		managers.NewDefaultDeployerManager,
		wire.Bind(new(managers.DeployerManager), new(managers.DefaultDeployerManager)),
	)

	return managers.DefaultDeployerManager{}, nil
}
