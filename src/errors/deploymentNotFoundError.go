package errors

import (
	"fmt"
	"github.com/aljrubior/anyctl/managers/entities"
)

func NewDeploymentNotFoundError(deploymentName string, options *[]entities.DeploymentItemEntity) *DeploymentNotFoundError {
	return &DeploymentNotFoundError{
		deploymentName,
		options,
	}
}

type DeploymentNotFoundError struct {
	DeploymentName string
	Options        *[]entities.DeploymentItemEntity
}

func (this *DeploymentNotFoundError) Error() string {
	return fmt.Sprintf("Deployment '%s' not found", this.DeploymentName)
}
