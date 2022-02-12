package wrappers

import (
	"github.com/aljrubior/anyctl/managers/entities"
	"github.com/aljrubior/anyctl/services"
)

func NewOrganizationWrapper(accountService services.AccountService, token, orgId string) *OrganizationWrapper {
	return &OrganizationWrapper{
		accountService: accountService,
		token:          token,
		orgId:          orgId,
	}

}

type OrganizationWrapper struct {
	accountService services.AccountService
	token          string
	orgId          string

	visitedOrgIds map[string]*entities.OrganizationEntity
}

func (this *OrganizationWrapper) FindMasterOrg() (*entities.OrganizationEntity, error) {

	this.visitedOrgIds = make(map[string]*entities.OrganizationEntity)

	resp, err := this.accountService.GetOrganization(this.token, this.orgId)

	if err != nil {
		return nil, err
	}

	if resp.IsMaster {
		return entities.NewOrganizationEntityBuilder(resp).Build(), nil
	}

	masterOrg, err := this.findMasterByParentIds(resp.ParentOrganizationIds)

	if err != nil {
		return nil, err
	}

	if masterOrg != nil {
		return masterOrg, nil
	}

	return nil, nil
}

func (this *OrganizationWrapper) FindAllOrgs() (*map[string]*entities.OrganizationEntity, error) {

	masterOrg, err := this.FindMasterOrg()

	if err != nil {
		return nil, err
	}

	if masterOrg == nil {
		return nil, nil
	}

	this.appendVisitedOrg(masterOrg)

	this.findSubOrgs(masterOrg.SubOrganizationIds)

	return &this.visitedOrgIds, nil
}

func (this *OrganizationWrapper) FindSingleOrg() (*entities.OrganizationEntity, error) {

	resp, err := this.accountService.GetOrganization(this.token, this.orgId)

	if err != nil {
		return nil, err
	}

	return entities.NewOrganizationEntityBuilder(resp).Build(), nil
}

func (this *OrganizationWrapper) findMasterByParentIds(fromParentOrgIds []string) (*entities.OrganizationEntity, error) {

	for _, v := range fromParentOrgIds {

		if _, found := this.visitedOrgIds[v]; found {
			continue
		}

		org, err := this.findMasterByOrgId(v)

		this.appendVisitedOrg(org)

		if err != nil {
			return nil, err
		}

		if org != nil {
			return org, nil
		}

		this.findMasterByParentIds(org.ParentOrganizationIds)
	}

	return nil, nil
}

func (this *OrganizationWrapper) findSubOrgs(subOrgs []string) (*entities.OrganizationEntity, error) {

	for _, v := range subOrgs {

		if _, found := this.visitedOrgIds[v]; found {
			continue
		}

		org, err := this.GetOrganization(v)

		this.appendVisitedOrg(org)

		if err != nil {
			return nil, err
		}

		this.findSubOrgs(org.SubOrganizationIds)
	}

	return nil, nil
}

func (this *OrganizationWrapper) findMasterByOrgId(orgId string) (*entities.OrganizationEntity, error) {

	resp, err := this.accountService.GetOrganization(this.token, orgId)

	if err != nil {
		return nil, err
	}

	if resp.IsMaster {
		return entities.NewOrganizationEntityBuilder(resp).Build(), nil
	}

	return nil, nil
}

func (this *OrganizationWrapper) GetOrganization(orgId string) (*entities.OrganizationEntity, error) {

	resp, err := this.accountService.GetOrganization(this.token, orgId)

	if err != nil {
		return nil, err
	}

	return entities.NewOrganizationEntityBuilder(resp).Build(), nil
}

func (this *OrganizationWrapper) appendVisitedOrg(org *entities.OrganizationEntity) {

	this.visitedOrgIds[org.Id] = org
}
