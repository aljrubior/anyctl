package handlers

import (
	errors2 "github.com/aljrubior/anyctl/errors"
	"github.com/aljrubior/anyctl/managers"
	"github.com/aljrubior/anyctl/managers/entities"
	"github.com/aljrubior/anyctl/manifests"
	"github.com/aljrubior/anyctl/printers"
	"github.com/aljrubior/anyctl/utils"
)

func NewDefaultFabricHandler(
	configManager managers.ConfigManager,
	fabricsManager managers.FabricManager,
	accountManager managers.AccountManager) *DefaultFabricHandler {

	return &DefaultFabricHandler{
		configManager,
		fabricsManager,
		accountManager,
	}
}

type DefaultFabricHandler struct {
	configManager  managers.ConfigManager
	fabricsManager managers.FabricManager
	accountManager managers.AccountManager
}

func (this DefaultFabricHandler) GetFabrics(fabricId string) error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	fabrics, err := this.fabricsManager.FindFabricByNameOrId(ctx, fabricId)

	if err != nil {
		return err
	}

	organizations := make(map[string]*entities.OrganizationEntity)

	for _, v := range *fabrics {
		org, err := this.accountManager.FindSingleOrg(ctx, v.OrganizationId)

		if err != nil {
			return err
		}

		organizations[org.Id] = org
	}

	utils.PrintFabrics(fabrics, organizations)

	return nil
}

func (this DefaultFabricHandler) GetFabric(fabricId string) error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	fabric, err := this.fabricsManager.GetFabric(ctx, fabricId)

	if err != nil {
		return err
	}

	if fabric == nil {
		return this.ThrowNewFabricNotFoundError(fabricId)
	}

	utils.PrintFabric(fabric)

	return nil
}

func (this DefaultFabricHandler) DescribeFabric(privateSpaceId string) error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	fabrics, err := this.fabricsManager.FindFabricByNameOrId(ctx, privateSpaceId)

	if err != nil {
		return err
	}

	if len(*fabrics) != 1 {
		return this.ThrowNewFabricNotFoundError(privateSpaceId)
	}

	utils.PrintFabricManifest(manifests.NewFabricManifest((*fabrics)[0].FabricResponse))

	return nil
}

func (this DefaultFabricHandler) GetVersions(fabricId string) error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	fabrics, err := this.fabricsManager.FindFabricByNameOrId(ctx, fabricId)

	if err != nil {
		return err
	}

	if len(*fabrics) != 1 {
		return this.ThrowNewFabricNotFoundError(fabricId)
	}

	printers.NewFabricPrinter(&(*fabrics)[0]).PrintVersionInformation()

	return nil
}

func (this DefaultFabricHandler) ThrowNewFabricNotFoundError(fabricId string) error {
	return errors2.NewFabricNotFoundError(fabricId, nil)
}
