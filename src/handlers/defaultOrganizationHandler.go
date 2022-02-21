package handlers

import (
	"github.com/aljrubior/anyctl/managers"
	"github.com/aljrubior/anyctl/managers/entities"
	"github.com/aljrubior/anyctl/printers"
)

func NewDefaultOrganizationHandler(accountManager managers.AccountManager, configManager managers.ConfigManager) *DefaultOrganizationHandler {
	return &DefaultOrganizationHandler{
		accountManager,
		configManager,
	}
}

type DefaultOrganizationHandler struct {
	accountManager managers.AccountManager
	configManager  managers.ConfigManager
}

func (this DefaultOrganizationHandler) GetCurrentOrganizationUsage() error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	org, err := this.accountManager.FindSingleOrg(ctx, ctx.OrganizationId)

	if err != nil {
		return err
	}

	printers.NewOrganizationUsagePrinter([]*entities.OrganizationEntity{org}).Print()

	return nil
}

func (this DefaultOrganizationHandler) GetAllOrganizationsUsage() error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	orgs, err := this.accountManager.FindAllOrgs(ctx, ctx.OrganizationId)

	if err != nil {
		return err
	}

	var orgsAsArray []*entities.OrganizationEntity

	for _, v := range *orgs {
		orgsAsArray = append(orgsAsArray, v)
	}

	printers.NewOrganizationUsagePrinter(orgsAsArray).Print()

	return nil
}

func (this DefaultOrganizationHandler) GetSingleOrganizationUsage(orgId string) error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	org, err := this.accountManager.FindSingleOrg(ctx, orgId)

	if err != nil {
		return err
	}

	printers.NewOrganizationUsagePrinter([]*entities.OrganizationEntity{org}).Print()

	return nil
}

func (this DefaultOrganizationHandler) GetCurrentOrganizationQuotas() error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	org, err := this.accountManager.FindSingleOrg(ctx, ctx.OrganizationId)

	if err != nil {
		return err
	}

	printers.NewOrganizationQuotasPrinter(org).Print()

	return nil
}
