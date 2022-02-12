package handlers

import (
	"fmt"
	errors2 "github.com/aljrubior/anyctl/errors"
	"github.com/aljrubior/anyctl/managers"
	"github.com/aljrubior/anyctl/managers/entities"
	"github.com/aljrubior/anyctl/utils"
)

func NewDefaultDeploymentMigrationHandler(
	deployerManager managers.DeployerManager,
	configManager managers.ConfigManager,
	deploymentManager managers.DeploymentManager,
	targetManager managers.TargetManager,
	accountManager managers.AccountManager) *DefaultDeploymentMigrationHandler {
	return &DefaultDeploymentMigrationHandler{
		configManager:     configManager,
		deployerManager:   deployerManager,
		deploymentManager: deploymentManager,
		targetManager:     targetManager,
		accountManager:    accountManager,
	}

}

type DefaultDeploymentMigrationHandler struct {
	configManager     managers.ConfigManager
	deployerManager   managers.DeployerManager
	deploymentManager managers.DeploymentManager
	targetManager     managers.TargetManager
	accountManager    managers.AccountManager
}

func (this DefaultDeploymentMigrationHandler) Migrate(fromDeploymentName, withName, toTargetName, toEnvironmentName string) error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	org, err := this.accountManager.FindSingleOrg(ctx, ctx.OrganizationId)

	if err != nil {
		return err
	}

	var toEnvironmentId string

	if toEnvironmentName == "" {
		toEnvironmentId = ctx.EnvironmentId
	} else {

		toEnvironmentId = this.findEnvironmentId(org, toEnvironmentName)

		if toEnvironmentId == "" {
			return errors2.NewEnvironmentNotFoundError(toEnvironmentName, org)
		}
	}

	fromDeployment, deployments, err := this.deploymentManager.FindDeploymentByName(ctx, fromDeploymentName)

	if err != nil {
		return err
	}

	if fromDeployment == nil {
		return this.ThrowDeploymentNotFoundError(fromDeploymentName, deployments)
	}

	var toTarget *entities.TargetEntity

	if toTargetName != "" {
		targetFound, options, err := this.targetManager.FindTargetByName(ctx, toTargetName)

		if err != nil {
			return err
		}

		if targetFound == nil {
			return this.ThrowTargetNotFoundError(toTargetName, options)
		}

		toTarget = targetFound
	}

	result, err := this.deployerManager.CopyDeployment(ctx, fromDeployment, withName, toTarget, toEnvironmentId)

	if err != nil {
		return err
	}

	if toEnvironmentName == "" {
		println(fmt.Sprintf("Deployment '%s' created on target '%s'.", result.Name, (*toTarget).GetName()))
	} else {
		println(fmt.Sprintf("Deployment '%s' created on target '%s' in the environment '%s'.", result.Name, (*toTarget).GetName(), toEnvironmentName))
	}

	this.deploymentManager.DeleteDeployment(ctx, fromDeployment.Id)

	targets, err := this.targetManager.GetTargets(ctx)

	targetsMap := utils.TargetEntities2Map(*targets)

	println(fmt.Sprintf("Deployment '%s' deleted on target '%s'.", fromDeployment.Name, targetsMap[fromDeployment.Target.TargetId].GetName()))

	return nil
}

func (this DefaultDeploymentMigrationHandler) Clone(fromDeploymentName, withName, toTargetName, toEnvironmentName string) error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	org, err := this.accountManager.FindSingleOrg(ctx, ctx.OrganizationId)

	if err != nil {
		return err
	}

	var toEnvironmentId string

	if toEnvironmentName == "" {
		toEnvironmentId = ctx.EnvironmentId
	} else {

		toEnvironmentId = this.findEnvironmentId(org, toEnvironmentName)

		if toEnvironmentId == "" {
			return errors2.NewEnvironmentNotFoundError(toEnvironmentName, org)
		}
	}

	fromDeployment, deployments, err := this.deploymentManager.FindDeploymentByName(ctx, fromDeploymentName)

	if err != nil {
		return err
	}

	if fromDeployment == nil {
		return this.ThrowDeploymentNotFoundError(fromDeploymentName, deployments)
	}

	var toTarget *entities.TargetEntity

	if toTargetName != "" {

		targetFound, options, err := this.targetManager.FindTargetByName(ctx, toTargetName)

		if err != nil {
			return err
		}

		if targetFound == nil {
			return this.ThrowTargetNotFoundError(toTargetName, options)
		}

		toTarget = targetFound
	}

	result, err := this.deployerManager.CopyDeployment(ctx, fromDeployment, withName, toTarget, toEnvironmentId)

	if err != nil {
		return err
	}

	if toTargetName == "" && toEnvironmentName == "" {
		println(fmt.Sprintf("Deployment '%s' created.", result.Name))
		return nil
	}

	if toEnvironmentName == "" {
		println(fmt.Sprintf("Deployment '%s' created on target '%s'.", result.Name, (*toTarget).GetName()))
		return nil
	}

	println(fmt.Sprintf("Deployment '%s' created on target '%s' in the environment '%s'.", result.Name, (*toTarget).GetName(), toEnvironmentName))

	return nil
}

func (this DefaultDeploymentMigrationHandler) ThrowDeploymentNotFoundError(deploymentName string, deployments *[]entities.DeploymentItemEntity) error {
	return errors2.NewDeploymentNotFoundError(deploymentName, deployments)
}

func (this DefaultDeploymentMigrationHandler) ThrowEnvironmentNotFoundError(deploymentName string, deployments *[]entities.DeploymentItemEntity) error {
	return errors2.NewDeploymentNotFoundError(deploymentName, deployments)
}

func (this DefaultDeploymentMigrationHandler) ThrowTargetNotFoundError(targetName string, options *[]entities.TargetEntity) error {
	return errors2.NewTargetNotFoundError(targetName, options)
}

func (this DefaultDeploymentMigrationHandler) findEnvironmentId(org *entities.OrganizationEntity, fromName string) string {
	for _, v := range org.Environments {
		if v.Name == fromName {
			return v.Id
		}
	}

	return ""
}
