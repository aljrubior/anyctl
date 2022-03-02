package handlers

import "github.com/aljrubior/anyctl/managers/entities"

type DeploymentMigrationHandler interface {
	Migrate(fromDeploymentName, withName, toTargetName, toEnvironment string) error
	Clone(fromDeploymentName, withName, toTargetName, toEnvironmentName string) error
	ThrowDeploymentNotFoundError(deploymentName string, deployments *[]entities.DeploymentItemEntity) error
}
