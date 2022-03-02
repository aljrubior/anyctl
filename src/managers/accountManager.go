package managers

import "github.com/aljrubior/anyctl/managers/entities"

type AccountManager interface {
	Login(username, password string) (*entities.LoginEntity, error)
	RefreshAccessToken(username, password string) (string, error)
	FindMasterOrg(ctx *entities.CurrentContextEntity, organizationId string) (*entities.OrganizationEntity, error)
	FindAllOrgs(ctx *entities.CurrentContextEntity, organizationId string) (*map[string]*entities.OrganizationEntity, error)
	FindSingleOrg(ctx *entities.CurrentContextEntity, orgId string) (*entities.OrganizationEntity, error)
}
