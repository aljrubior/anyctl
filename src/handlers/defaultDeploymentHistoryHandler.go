package handlers

import (
	"github.com/aljrubior/anyctl/comparators"
	errors2 "github.com/aljrubior/anyctl/errors"
	"github.com/aljrubior/anyctl/managers"
	"github.com/aljrubior/anyctl/managers/entities"
	"github.com/aljrubior/anyctl/printers"
)

func NewDefaultDeploymentHistoryHandler(
	deploymentManager managers.DeploymentManager,
	configManager managers.ConfigManager) DefaultDeploymentHistoryHandler {

	return DefaultDeploymentHistoryHandler{
		deploymentManager: deploymentManager,
		configManager:     configManager,
	}

}

type DefaultDeploymentHistoryHandler struct {
	deploymentManager managers.DeploymentManager
	configManager     managers.ConfigManager
}

func (this DefaultDeploymentHistoryHandler) GetDeploymentHistory(deploymentName string) error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	deployment, deployments, err := this.deploymentManager.FindDeploymentByName(ctx, deploymentName)

	if err != nil {
		return err
	}

	if deployment == nil {
		return this.ThrowDeploymentNotFoundError(deploymentName, deployments)
	}

	specs, err := this.deploymentManager.GetDeploymentSpecs(ctx, deployment.Id)

	if err != nil {
		return err
	}

	printers.NewDeploymentSpecPrinter(specs).Print()

	return nil
}

func (this DefaultDeploymentHistoryHandler) Compare(deploymentName, withSpecVersion string) error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	deployment, deployments, err := this.deploymentManager.FindDeploymentByName(ctx, deploymentName)

	if err != nil {
		return err
	}

	if deployment == nil {
		return this.ThrowDeploymentNotFoundError(deploymentName, deployments)
	}

	specs, err := this.deploymentManager.GetDeploymentSpecs(ctx, deployment.Id)

	if err != nil {
		return err
	}

	var currentVersion, withSpec entities.DeploymentSpecEntity

	for i, k := range *specs {
		if i == 0 {
			currentVersion = k
			continue
		}

		if k.Version[:6] == withSpecVersion {
			withSpec = k
			break
		}
	}

	comparator, err := comparators.NewDeploymentSpecResponseComparator(currentVersion.DeploymentSpecResponse, withSpec.DeploymentSpecResponse)

	if err != nil {
		return err
	}

	differences := comparator.Compare()

	printers.NewDeploymentDifferencePrinter(differences).Print()

	return nil
}

func (this DefaultDeploymentHistoryHandler) ThrowDeploymentNotFoundError(deploymentName string, deployments *[]entities.DeploymentItemEntity) error {
	return errors2.NewDeploymentNotFoundError(deploymentName, deployments)
}
