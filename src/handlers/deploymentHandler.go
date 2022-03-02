package handlers

import "github.com/aljrubior/anyctl/managers/entities"

type DeploymentHandler interface {
	GetDeployments() error
	GetDeployment(deploymentName string) error
	StopDeployment(deploymentName string) error
	StartDeployment(deploymentName string) error
	DeleteDeployment(deploymentName string) error
	ScaleDeployment(deploymentName string, desiredReplicas int) error
	SetDeploymentAsset(deploymentName string, assetRef string) error
	FindDeploymentsContainsName(deploymentName string) error
	DescribeDeployment(deploymentName string) error

	ThrowDeploymentNotFoundError(deploymentName string, deployments *[]entities.DeploymentItemEntity) error
}
