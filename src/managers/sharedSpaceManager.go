package managers

import "github.com/aljrubior/anyctl/managers/entities"

type SharedSpaceManager interface {
	GetSharedSpace(ctx *entities.CurrentContextEntity, sharedSapceId string) (*entities.SharedSpaceEntity, error)
	GetSharedSpaces(ctx *entities.CurrentContextEntity) (*[]entities.SharedSpaceEntity, error)

	FindSharedSpaceByName(ctx *entities.CurrentContextEntity, sharedSpaceName string) (*entities.SharedSpaceEntity, *[]entities.SharedSpaceEntity, error)
	FindSharedSpaceContainsName(ctx *entities.CurrentContextEntity, sharedSpaceName string) (*[]entities.SharedSpaceEntity, error)
}
