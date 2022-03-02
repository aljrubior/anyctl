package managers

import (
	"github.com/aljrubior/anyctl/managers/entities"
	"github.com/aljrubior/anyctl/managers/requests"
	"github.com/aljrubior/anyctl/managers/wrappers"
	"github.com/aljrubior/anyctl/services"
)

func NewDefaultAccountManager(accountService services.AccountService) DefaultAccountManager {
	return DefaultAccountManager{
		accountService,
	}
}

type DefaultAccountManager struct {
	accountService services.AccountService
}

func (this DefaultAccountManager) Login(username, password string) (*entities.LoginEntity, error) {

	req := requests.NewLoginRequestBuilder(username, password).Build()

	token, err := this.accountService.Login(*req)

	if err != nil {
		return nil, err
	}

	profile, err := this.accountService.GetProfile(*token)

	if err != nil {
		return nil, err
	}

	return entities.NewLoginEntityBuilder(*token, profile).Build(), nil
}

func (this DefaultAccountManager) RefreshAccessToken(username, password string) (string, error) {

	loginEntity, err := this.Login(username, password)

	if err != nil {
		return "", err
	}

	return loginEntity.Token, nil
}

func (this DefaultAccountManager) FindMasterOrg(ctx *entities.CurrentContextEntity, organizationId string) (*entities.OrganizationEntity, error) {

	wrapper := wrappers.NewOrganizationWrapper(this.accountService, ctx.AuthorizationToken, organizationId)

	masterOrg, err := wrapper.FindMasterOrg()

	if err != nil {
		return nil, err
	}

	if masterOrg != nil {
		return masterOrg, nil
	}

	return nil, nil
}

func (this DefaultAccountManager) FindAllOrgs(ctx *entities.CurrentContextEntity, organizationId string) (*map[string]*entities.OrganizationEntity, error) {

	wrapper := wrappers.NewOrganizationWrapper(this.accountService, ctx.AuthorizationToken, organizationId)

	orgs, err := wrapper.FindAllOrgs()

	if err != nil {
		return nil, err
	}

	return orgs, nil
}

func (this DefaultAccountManager) FindSingleOrg(ctx *entities.CurrentContextEntity, orgId string) (*entities.OrganizationEntity, error) {

	wrapper := wrappers.NewOrganizationWrapper(this.accountService, ctx.AuthorizationToken, orgId)

	return wrapper.FindSingleOrg()
}
