package handlers

import (
	"fmt"
	errors2 "github.com/aljrubior/anyctl/errors"
	"github.com/aljrubior/anyctl/managers"
	"github.com/aljrubior/anyctl/managers/entities"
	"github.com/aljrubior/anyctl/manifests"
	"github.com/aljrubior/anyctl/printers"
	"github.com/aljrubior/anyctl/utils"
)

func NewDeploymentHandler(deploymentManager managers.DeploymentManager, configManager managers.ConfigManager, targetManager managers.TargetManager) DefaultDeploymentHandler {
	return DefaultDeploymentHandler{
		deploymentManager: deploymentManager,
		configManager:     configManager,
		targetManager:     targetManager,
	}
}

type DefaultDeploymentHandler struct {
	deploymentManager managers.DeploymentManager
	configManager     managers.ConfigManager
	targetManager     managers.TargetManager
}

func (this DefaultDeploymentHandler) GetDeployments() error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	deployments, err := this.deploymentManager.GetDeployments(ctx)

	targets, err := this.targetManager.GetTargets(ctx)

	if err != nil {
		return err
	}

	utils.PrintDeployments(deployments, targets)

	return nil
}

func (this DefaultDeploymentHandler) GetDeployment(deploymentName string) error {

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

	targets, err := this.targetManager.GetTargets(ctx)

	if err != nil {
		return err
	}

	utils.PrintDeployment(deployment, targets)

	return nil
}

func (this DefaultDeploymentHandler) FindDeploymentsContainsName(deploymentName string) error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	deployments, err := this.deploymentManager.FindDeploymentContainsName(ctx, deploymentName)

	if err != nil {
		return err
	}

	targets, err := this.targetManager.GetTargets(ctx)

	if err != nil {
		return err
	}

	if len(*deployments) == 0 {
		return this.ThrowDeploymentNotFoundError(deploymentName, nil)
	}

	if len(*deployments) == 1 {
		deployment, err := this.deploymentManager.GetDeployment(ctx, (*deployments)[0].Id)

		if err != nil {
			return nil
		}

		utils.PrintDeployment(deployment, targets)
		return nil
	}

	utils.PrintDeployments(deployments, targets)

	return nil
}

func (this DefaultDeploymentHandler) StopDeployment(deploymentName string) error {

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

	deployment, err = this.deploymentManager.StopDeployment(ctx, deployment.Id)

	if err != nil {
		return err
	}

	println(fmt.Sprintf("Deployment '%s' stopped.", deployment.Name))

	return nil
}

func (this DefaultDeploymentHandler) StartDeployment(deploymentName string) error {

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

	deployment, err = this.deploymentManager.StartDeployment(ctx, deployment.Id)

	if err != nil {
		return err
	}

	println(fmt.Sprintf("Deployment '%s' started.", deployment.Name))

	return nil
}

func (this DefaultDeploymentHandler) DeleteDeployment(deploymentName string) error {

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

	err = this.deploymentManager.DeleteDeployment(ctx, deployment.Id)

	if err != nil {
		return err
	}

	println(fmt.Sprintf("Deployment '%s' deleted.", deploymentName))

	return nil
}

func (this DefaultDeploymentHandler) ScaleDeployment(deploymentName string, desiredReplicas int) error {

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

	deployment, err = this.deploymentManager.ScaleDeployment(ctx, deployment.Id, desiredReplicas)

	if err != nil {
		return err
	}

	println(fmt.Sprintf("Deployment '%s' scaled to '%d' replicas.", deployment.Name, deployment.Target.Replicas))

	return nil
}

func (this DefaultDeploymentHandler) SetDeploymentAsset(deploymentName string, assetRef string) error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	deployment, err := this.deploymentManager.SetDeploymentAsset(ctx, deploymentName, assetRef)

	if err != nil {
		return err
	}

	println(fmt.Sprintf("Deployment '%s' updated with asset '%s:%s'.", deployment.Name, deployment.Application.Ref.ArtifactId, deployment.Application.Ref.Version))

	return nil
}

func (this DefaultDeploymentHandler) DescribeDeployment(deploymentName string) error {

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

	printer, err := printers.NewDeploymentManifestPrinter(manifests.NewDeploymentManifest(deployment.DeploymentResponse))

	if err != nil {
		return err
	}

	printer.Print()

	return nil
}

func (this DefaultDeploymentHandler) ThrowDeploymentNotFoundError(deploymentName string, deployments *[]entities.DeploymentItemEntity) error {
	return errors2.NewDeploymentNotFoundError(deploymentName, deployments)
}
