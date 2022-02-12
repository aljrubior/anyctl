package handlers

import (
	"github.com/aljrubior/anyctl/managers"
	"github.com/aljrubior/anyctl/utils"
)

func NewDefaultEntitlementHandler(accountManager managers.AccountManager, configManager managers.ConfigManager) *DefaultEntitlementHandler {
	return &DefaultEntitlementHandler{
		accountManager,
		configManager,
	}
}

type DefaultEntitlementHandler struct {
	accountManager managers.AccountManager
	configManager  managers.ConfigManager
}

func (this DefaultEntitlementHandler) GetCurrentOrganizationEntitlement() error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	org, err := this.accountManager.FindSingleOrg(ctx, ctx.OrganizationId)

	if err != nil {
		return err
	}

	utils.PrintOrgQuotas(org)

	return nil
}
