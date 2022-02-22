package handlers

import (
	"github.com/aljrubior/anyctl/errors"
	"github.com/aljrubior/anyctl/managers"
	"github.com/aljrubior/anyctl/manifests"
	"github.com/aljrubior/anyctl/printers"
	"github.com/aljrubior/anyctl/utils"
)

func NewDefaultPrivateSpaceHandler(configManager managers.ConfigManager, privateSpaceManager managers.PrivateSpaceManager) *DefaultPrivateSpaceHandler {
	return &DefaultPrivateSpaceHandler{
		configManager,
		privateSpaceManager,
	}
}

type DefaultPrivateSpaceHandler struct {
	configManager       managers.ConfigManager
	privateSpaceManager managers.PrivateSpaceManager
}

func (this DefaultPrivateSpaceHandler) GetPrivateSpaces(privateSpaceId string) error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	privateSpaces, err := this.privateSpaceManager.FindPrivateSpaceByNameOrId(ctx, privateSpaceId)

	if err != nil {
		return err
	}

	utils.PrintPrivateSpaces(privateSpaces)

	return nil
}

func (this DefaultPrivateSpaceHandler) GetPrivateSpace(privateSpaceId string) error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	privateSpace, err := this.privateSpaceManager.GetPrivateSpace(ctx, privateSpaceId)

	if err != nil {
		return err
	}

	if privateSpace == nil {
		return this.ThrowNewPrivateSpaceNotFoundError(privateSpaceId)
	}

	utils.PrintPrivateSpace(privateSpace)

	return nil
}

func (this DefaultPrivateSpaceHandler) GetManagedFirewallRules(privateSpaceId string) error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	privateSpaces, err := this.privateSpaceManager.FindPrivateSpaceByNameOrId(ctx, privateSpaceId)

	if err != nil {
		return err
	}

	if len(*privateSpaces) != 1 {
		return this.ThrowNewPrivateSpaceNotFoundError(privateSpaceId)
	}

	printers.NewPrivateSpacePrinter(&(*privateSpaces)[0]).PrintManagedFirewallRules()

	return nil
}

func (this DefaultPrivateSpaceHandler) DescribePrivateSpace(privateSpaceId string) error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	privateSpaces, err := this.privateSpaceManager.FindPrivateSpaceByNameOrId(ctx, privateSpaceId)

	if err != nil {
		return err
	}

	if len(*privateSpaces) != 1 {
		return this.ThrowNewPrivateSpaceNotFoundError(privateSpaceId)
	}

	printer, err := printers.NewPrivateSpaceManifestPrinter(manifests.NewPrivateSpaceManifest((*privateSpaces)[0].PrivateSpaceResponse))

	if err != nil {
		return err
	}

	printer.Print()

	return nil
}

func (this DefaultPrivateSpaceHandler) GetFabrics(privateSpaceId string) error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	privateSpaces, err := this.privateSpaceManager.FindPrivateSpaceByNameOrId(ctx, privateSpaceId)

	if err != nil {
		return err
	}

	if len(*privateSpaces) != 1 {
		return this.ThrowNewPrivateSpaceNotFoundError(privateSpaceId)
	}

	fabrics, err := this.privateSpaceManager.GetFabrics(ctx, (*privateSpaces)[0].Id)

	if err != nil {
		return err
	}

	utils.PrintPrivateSpaceFabrics(fabrics)

	return nil
}

func (this DefaultPrivateSpaceHandler) ThrowNewPrivateSpaceNotFoundError(privateSpaceId string) error {
	return errors.NewPrivateSpaceNotFoundError(privateSpaceId, nil)
}
