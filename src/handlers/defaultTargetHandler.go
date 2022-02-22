package handlers

import (
	"github.com/aljrubior/anyctl/errors"
	"github.com/aljrubior/anyctl/managers"
	"github.com/aljrubior/anyctl/managers/entities"
	"github.com/aljrubior/anyctl/managers/wrappers"
	"github.com/aljrubior/anyctl/manifests"
	"github.com/aljrubior/anyctl/printers"
	"github.com/aljrubior/anyctl/utils"
	"strings"
)

func NewDefaultTargetHandler(configManager managers.ConfigManager, targetManager managers.TargetManager) *DefaultTargetHandler {
	return &DefaultTargetHandler{
		configManager: configManager,
		targetManager: targetManager,
	}
}

type DefaultTargetHandler struct {
	configManager managers.ConfigManager
	targetManager managers.TargetManager
}

func (this DefaultTargetHandler) GetTargets() error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	targets, err := this.targetManager.GetTargets(ctx)

	if err != nil {
		return err
	}

	printers.NewTargetsPrinter(targets).Print()

	return nil
}

func (this DefaultTargetHandler) GetTarget(targetName string) error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	targets, err := this.targetManager.GetTargets(ctx)

	if err != nil {
		return err
	}

	var targetsFound []entities.TargetEntity

	for _, v := range *targets {
		if strings.Contains(v.GetName(), targetName) {
			targetsFound = append(targetsFound, v)
		}
	}

	printers.NewTargetsPrinter(&targetsFound).Print()

	return nil
}

func (this DefaultTargetHandler) FindTargetsContainsName(targetName string) error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	targets, err := this.targetManager.FindTargetsContainsName(ctx, targetName)

	if err != nil {
		return err
	}

	printers.NewTargetsPrinter(targets).Print()

	return nil
}

func (this DefaultTargetHandler) GetSupportedRuntimes(targetName string) error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	target, options, err := this.targetManager.FindTargetByName(ctx, targetName)

	if err != nil {
		return err
	}

	if target == nil {
		return this.ThrowTargetNotFoundError(targetName, options)
	}

	printers.NewTargetPrinter(target).PrintTargetSupportedVersions()

	return nil
}

func (this DefaultTargetHandler) GetDetails(targetName string) error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	target, options, err := this.targetManager.FindTargetByName(ctx, targetName)

	if err != nil {
		return err
	}

	if target == nil {
		return this.ThrowTargetNotFoundError(targetName, options)
	}

	printers.NewTargetPrinter(target).PrintStandaloneDetails()

	return nil
}

func (this DefaultTargetHandler) GetAddresses(targetName string) error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	target, options, err := this.targetManager.FindTargetByName(ctx, targetName)

	if err != nil {
		return err
	}

	if target == nil {
		return this.ThrowTargetNotFoundError(targetName, options)
	}

	utils.PrintStandaloneAddresses(wrappers.NewTargetEntityWrapper(*target))

	return nil
}

func (this DefaultTargetHandler) DescribeTarget(targetName string) error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	target, options, err := this.targetManager.FindTargetByName(ctx, targetName)

	if err != nil {
		return err
	}

	if target == nil {
		return this.ThrowTargetNotFoundError(targetName, options)
	}

	targetWrapper := wrappers.NewTargetEntityWrapper(*target)

	if targetWrapper.IsStandaloneTargetEntity() {

		standaloneTarget, _ := targetWrapper.GetStandaloneTargetEntity()

		printer, err := printers.NewStandaloneTargetManifestPrinter(manifests.NewStandaloneTargetManifest(standaloneTarget.StandaloneTargetResponse))

		if err != nil {
			return nil
		}

		printer.Print()

		return nil
	}

	if targetWrapper.IsRuntimeFabricTargetEntity() {
		runtimeFabricTarget, _ := targetWrapper.GetRuntimeFabricTargetEntity()
		printer, err := printers.NewRuntimeFabricTargetManifestPrinter(manifests.NewRuntimeFabricTargetManifest(runtimeFabricTarget.RuntimeFabricTargetResponse))

		if err != nil {
			return err
		}

		printer.Print()
		return nil
	}

	return nil
}

func (this DefaultTargetHandler) ThrowTargetNotFoundError(targetName string, targets *[]entities.TargetEntity) error {
	return errors.NewTargetNotFoundError(targetName, targets)
}
