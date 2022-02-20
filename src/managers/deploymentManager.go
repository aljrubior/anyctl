package managers

import (
	"github.com/aljrubior/anyctl/clients/deployments/response"
	"github.com/aljrubior/anyctl/managers/entities"
	"github.com/aljrubior/anyctl/managers/requests"
)

type DeploymentManager interface {
	GetDeployments(ctx *entities.CurrentContextEntity) (*[]entities.DeploymentItemEntity, error)
	GetDeployment(ctx *entities.CurrentContextEntity, deploymentId string) (*entities.DeploymentEntity, error)
	FindDeploymentByName(ctx *entities.CurrentContextEntity, deploymentName string) (*entities.DeploymentEntity, *[]entities.DeploymentItemEntity, error)
	FindDeploymentContainsName(ctx *entities.CurrentContextEntity, deploymentName string) (*[]entities.DeploymentItemEntity, error)
	Deploy(ctx *entities.CurrentContextEntity, request *requests.DeploymentRequest, toEnvironment string) (*response.DeploymentResponse, error)
	StopDeployment(ctx *entities.CurrentContextEntity, deploymentId string) (*entities.DeploymentEntity, error)
	StartDeployment(ctx *entities.CurrentContextEntity, deploymentId string) (*entities.DeploymentEntity, error)
	DeleteDeployment(ctx *entities.CurrentContextEntity, deploymentId string) error
	ScaleDeployment(ctx *entities.CurrentContextEntity, deploymentId string, desiredReplicas int) (*entities.DeploymentEntity, error)
	SetDeploymentAsset(ctx *entities.CurrentContextEntity, deploymentName, assetRef string) (*entities.DeploymentEntity, error)

	GetDeploymentSpecs(ctx *entities.CurrentContextEntity, deploymentId string) (*[]entities.DeploymentSpecEntity, error)
}
