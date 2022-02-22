package handlers

import (
	"github.com/aljrubior/anyctl/errors"
	"github.com/aljrubior/anyctl/managers"
	"github.com/aljrubior/anyctl/managers/entities"
	"github.com/aljrubior/anyctl/manifests"
	"github.com/aljrubior/anyctl/printers"
	"github.com/aljrubior/anyctl/utils"
)

func NewDefaultSharedSpaceHandler(
	configManager managers.ConfigManager,
	sharedSpaceManager managers.SharedSpaceManager,
	privateSpaceManager managers.PrivateSpaceManager) *DefaultSharedSpaceHandler {

	return &DefaultSharedSpaceHandler{
		configManager,
		sharedSpaceManager,
		privateSpaceManager,
	}
}

type DefaultSharedSpaceHandler struct {
	configManager       managers.ConfigManager
	sharedSpaceManager  managers.SharedSpaceManager
	privateSpaceManager managers.PrivateSpaceManager
}

func (this DefaultSharedSpaceHandler) GetSharedSpaces() error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	sharedSpaces, err := this.sharedSpaceManager.GetSharedSpaces(ctx)

	if err != nil {
		return err
	}

	utils.PrintSharedSpaces(sharedSpaces)

	return nil
}

func (this DefaultSharedSpaceHandler) GetSharedSpace(sharedSpaceName string) error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	sharedSpace, options, err := this.sharedSpaceManager.FindSharedSpaceByName(ctx, sharedSpaceName)

	if err != nil {
		return err
	}

	if sharedSpace == nil {
		return this.ThrowNewSharedSpaceNotFoundError(sharedSpaceName, options)
	}

	privateSpace, err := this.privateSpaceManager.GetPrivateSpace(ctx, sharedSpace.PrivateSpaceId)

	if err != nil {
		return nil
	}

	printers.NewSharedSpacePrinter(sharedSpace, privateSpace).Print()

	return nil
}

func (this DefaultSharedSpaceHandler) FindSharedSpaceContainsName(sharedSpaceName string) error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	sharedSpaces, err := this.sharedSpaceManager.FindSharedSpaceContainsName(ctx, sharedSpaceName)

	if err != nil {
		return err
	}

	if len(*sharedSpaces) == 0 {
		return this.ThrowNewSharedSpaceNotFoundError(sharedSpaceName, nil)
	}

	if len(*sharedSpaces) == 1 {
		privateSpace, err := this.privateSpaceManager.GetPrivateSpace(ctx, (*sharedSpaces)[0].PrivateSpaceId)

		if err != nil {
			return nil
		}

		printers.NewSharedSpacePrinter(&(*sharedSpaces)[0], privateSpace).Print()
		return nil

	}

	utils.PrintSharedSpaces(sharedSpaces)
	return nil
}

func (this DefaultSharedSpaceHandler) DescribeSharedSpace(sharedSpaceName string) error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	foundSharedSpaced, options, err := this.sharedSpaceManager.FindSharedSpaceByName(ctx, sharedSpaceName)

	if err != nil {
		return err
	}

	if foundSharedSpaced == nil {
		return this.ThrowNewSharedSpaceNotFoundError(sharedSpaceName, options)
	}

	sharedSpace, err := this.sharedSpaceManager.GetSharedSpace(ctx, foundSharedSpaced.Id)

	if err != nil {
		return err
	}

	printer, err := printers.NewSharedSpaceManifestPrinter(manifests.NewSharedSpaceManifest(sharedSpace.SharedSpaceResponse))

	if err != nil {
		return err
	}

	printer.Print()

	return nil
}

func (this DefaultSharedSpaceHandler) ThrowNewSharedSpaceNotFoundError(ssName string, options *[]entities.SharedSpaceEntity) error {
	return errors.NewSharedSpaceNotFoundError(ssName, options)
}
