package managers

import (
	"github.com/aljrubior/anyctl/managers/entities"
)

type DeployerManager interface {
	Deploy(ctx *entities.CurrentContextEntity, deploymentName, assetRef, targetName, runtimeBaseVersion string) (*entities.DeploymentEntity, error)
	CopyDeployment(ctx *entities.CurrentContextEntity, fromDeployment *entities.DeploymentEntity, withName string, toTarget *entities.TargetEntity, toEnvironmentId string) (*entities.DeploymentEntity, error)

	ThrowAssetNotFoundError(assetName string, err error) error
	ThrowFabricTargetNotFoundError(targetName string, options *[]entities.FabricTargetEntity) error
}
