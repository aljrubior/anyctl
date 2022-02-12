package managers

import "github.com/aljrubior/anyctl/managers/entities"

type OrganizationPrivateSpaceManager interface {
	GetPrivateSpaces(ctx *entities.CurrentContextEntity) (*[]entities.OrganizationPrivateSpaceEntity, error)
	FindPrivateSpaceByName(ctx *entities.CurrentContextEntity, psName string) (*entities.OrganizationPrivateSpaceEntity, *[]entities.OrganizationPrivateSpaceEntity, error)
	FindPrivateSpaceContainsName(ctx *entities.CurrentContextEntity, privateSpaceName string) (*[]entities.OrganizationPrivateSpaceEntity, error)
	GetFabrics(ctx *entities.CurrentContextEntity, privateSpaceId string) (*[]entities.OrganizationPrivateSpaceFabricEntity, error)
}
