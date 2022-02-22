package handlers

import (
	"github.com/aljrubior/anyctl/errors"
	"github.com/aljrubior/anyctl/managers"
	"github.com/aljrubior/anyctl/managers/entities"
	"github.com/aljrubior/anyctl/manifests"
	"github.com/aljrubior/anyctl/printers"
	"github.com/aljrubior/anyctl/utils"
)

func NewDefaultOrganizationPrivateSpaceHandler(configManager managers.ConfigManager, privateSpaceManager managers.OrganizationPrivateSpaceManager) *DefaultOrganizationPrivateSpaceHandler {
	return &DefaultOrganizationPrivateSpaceHandler{
		configManager,
		privateSpaceManager,
	}
}

type DefaultOrganizationPrivateSpaceHandler struct {
	configManager       managers.ConfigManager
	privateSpaceManager managers.OrganizationPrivateSpaceManager
}

func (this DefaultOrganizationPrivateSpaceHandler) GetPrivateSpaces() error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	privateSpaces, err := this.privateSpaceManager.GetPrivateSpaces(ctx)

	if err != nil {
		return err
	}

	printers.NewOrganizationPrivateSpacesPrinter(privateSpaces).Print()

	return nil
}

func (this DefaultOrganizationPrivateSpaceHandler) GetPrivateSpace(privateSpaceName string) error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	privateSpace, options, err := this.privateSpaceManager.FindPrivateSpaceByName(ctx, privateSpaceName)

	if err != nil {
		return err
	}

	if privateSpace == nil {
		return this.ThrowNewPrivateSpaceNotFoundError(privateSpaceName, options)
	}

	printers.NewOrganizationPrivateSpacePrinter(privateSpace).Print()

	return nil
}

func (this DefaultOrganizationPrivateSpaceHandler) FindPrivateSpaceContainsName(targetName string) error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	targets, err := this.privateSpaceManager.FindPrivateSpaceContainsName(ctx, targetName)

	if err != nil {
		return err
	}

	if len(*targets) == 1 {
		printers.NewOrganizationPrivateSpacePrinter(&(*targets)[0]).Print()
		return nil
	}

	printers.NewOrganizationPrivateSpacesPrinter(targets).Print()

	return nil
}

func (this DefaultOrganizationPrivateSpaceHandler) GetFirewallRules(psName string) error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	privateSpace, options, err := this.privateSpaceManager.FindPrivateSpaceByName(ctx, psName)

	if err != nil {
		return err
	}

	if privateSpace == nil {
		return this.ThrowNewPrivateSpaceNotFoundError(psName, options)
	}

	printers.NewOrganizationPrivateSpacePrinter(privateSpace).PrintFirewallRules()

	return nil
}

func (this DefaultOrganizationPrivateSpaceHandler) GetNetwork(psName string) error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	privateSpace, options, err := this.privateSpaceManager.FindPrivateSpaceByName(ctx, psName)

	if err != nil {
		return err
	}

	if privateSpace == nil {
		return this.ThrowNewPrivateSpaceNotFoundError(psName, options)
	}

	printers.NewOrganizationPrivateSpacePrinter(privateSpace).PrintNetwork()

	return nil
}

func (this DefaultOrganizationPrivateSpaceHandler) GetFabrics(privateSpaceName string) error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	privateSpace, options, err := this.privateSpaceManager.FindPrivateSpaceByName(ctx, privateSpaceName)

	if err != nil {
		return nil
	}

	if privateSpace == nil {
		return this.ThrowNewPrivateSpaceNotFoundError(privateSpaceName, options)
	}

	fabrics, err := this.privateSpaceManager.GetFabrics(ctx, privateSpace.Id)

	if err != nil {
		return err
	}

	utils.PrintOrganizationPrivateSpaceFabrics(fabrics)

	return nil
}

func (this DefaultOrganizationPrivateSpaceHandler) DescribePrivateSpace(psName string) error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	privateSpace, options, err := this.privateSpaceManager.FindPrivateSpaceByName(ctx, psName)

	if err != nil {
		return err
	}

	if privateSpace == nil {
		return this.ThrowNewPrivateSpaceNotFoundError(psName, options)
	}

	printer, err := printers.NewOrganizationPrivateSpaceManifestPrinter(manifests.NewOrganizationPrivateSpaceManifest(privateSpace.OrganizationPrivateSpaceResponse))

	if err != nil {
		return nil
	}

	printer.Print()

	return nil
}

func (this DefaultOrganizationPrivateSpaceHandler) ThrowNewPrivateSpaceNotFoundError(privateSpaceName string, options *[]entities.OrganizationPrivateSpaceEntity) error {
	return errors.NewOrganizationPrivateSpaceNotFoundError(privateSpaceName, options)
}
