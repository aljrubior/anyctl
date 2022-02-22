package handlers

import (
	"github.com/aljrubior/anyctl/errors"
	"github.com/aljrubior/anyctl/managers"
	"github.com/aljrubior/anyctl/managers/entities"
	"github.com/aljrubior/anyctl/manifests"
	"github.com/aljrubior/anyctl/printers"
	"github.com/aljrubior/anyctl/utils"
)

func NewDefaultRuntimeFabricHandler(configManager managers.ConfigManager, runtimeFabricManager managers.OrganizationRuntimeFabricManager) *DefaultRuntimeFabrichandler {
	return &DefaultRuntimeFabrichandler{
		configManager,
		runtimeFabricManager,
	}
}

type DefaultRuntimeFabrichandler struct {
	configManager        managers.ConfigManager
	runtimeFabricManager managers.OrganizationRuntimeFabricManager
}

func (this DefaultRuntimeFabrichandler) GetFabrics() error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	targets, err := this.runtimeFabricManager.GetFabrics(ctx)

	if err != nil {
		return err
	}

	printers.NewOrganizationFabricsEntity(targets).Print()

	return nil
}

func (this DefaultRuntimeFabrichandler) GetFabric(targetName string) error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	target, targets, err := this.runtimeFabricManager.FindFabricByName(ctx, targetName)

	if err != nil {
		return err
	}

	if target == nil {
		return this.ThrowNewOrganizationFabricNotFoundError(targetName, targets)
	}

	utils.PrintOrganizationFabric(target)

	return nil
}

func (this DefaultRuntimeFabrichandler) FindRuntimeFabricContainsName(targetName string) error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	targets, err := this.runtimeFabricManager.FindRuntimeFabricContainsName(ctx, targetName)

	if err != nil {
		return err
	}

	if len(*targets) == 1 {
		utils.PrintOrganizationFabric(&(*targets)[0])
		return nil
	}

	printers.NewOrganizationFabricsEntity(targets).Print()

	return nil
}

func (this DefaultRuntimeFabrichandler) DescribeFabric(targetName string) error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	target, targets, err := this.runtimeFabricManager.FindFabricByName(ctx, targetName)

	if err != nil {
		return err
	}

	if target == nil {
		return this.ThrowNewOrganizationFabricNotFoundError(targetName, targets)
	}

	utils.PrintRuntimeFabricManifest(manifests.NewOrganizationFabricManifest(target.OrganizationFabricResponse))

	return nil
}

func (this DefaultRuntimeFabrichandler) GetRuntimeFabricNodes(targetName string) error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	target, targets, err := this.runtimeFabricManager.FindFabricByName(ctx, targetName)

	if err != nil {
		return err
	}

	if target == nil {
		return this.ThrowNewOrganizationFabricNotFoundError(targetName, targets)
	}

	utils.PrintOrganzationFabricNodes(target)

	return nil
}

func (this DefaultRuntimeFabrichandler) ThrowNewOrganizationFabricNotFoundError(targetName string, targets *[]entities.OrganizationFabricEntity) error {
	return errors.NewOrganizationFabricNotFoundError(targetName, targets)
}
