package handlers

import (
	errors2 "github.com/aljrubior/anyctl/errors"
	"github.com/aljrubior/anyctl/managers"
	"github.com/aljrubior/anyctl/managers/entities"
	"github.com/aljrubior/anyctl/printers"
)

func NewDefaultDeploymentLogsHandler(
	deploymentManager managers.DeploymentManager,
	configManager managers.ConfigManager,
	deploymentLogsManager managers.DeploymentLogsManager) DefaultDeploymentLogsHandler {

	return DefaultDeploymentLogsHandler{
		deploymentManager:     deploymentManager,
		configManager:         configManager,
		deploymentLogsManager: deploymentLogsManager,
	}
}

type DefaultDeploymentLogsHandler struct {
	deploymentManager     managers.DeploymentManager
	configManager         managers.ConfigManager
	deploymentLogsManager managers.DeploymentLogsManager
}

func (this DefaultDeploymentLogsHandler) GetLogs(deploymentName string) error {

	ctx, err := this.configManager.GetCurrentContext()

	deployment, options, err := this.deploymentManager.FindDeploymentByName(ctx, deploymentName)

	if err != nil {
		return err
	}

	if deployment == nil {
		return this.ThrowDeploymentNotFoundError(deploymentName, options)
	}

	specs, err := this.deploymentManager.GetDeploymentSpecs(ctx, deployment.Id)

	if err != nil {
		return err
	}

	logs, err := this.deploymentLogsManager.GetLogs(ctx, deployment.Id, (*specs)[0].Version)

	if err != nil {
		return err
	}

	printers.NewDeploymentLogsPrinter(logs).Print()

	return nil
}

func (this DefaultDeploymentLogsHandler) ThrowDeploymentNotFoundError(deploymentName string, options *[]entities.DeploymentItemEntity) error {
	return errors2.NewDeploymentNotFoundError(deploymentName, options)
}
