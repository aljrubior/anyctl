package managers

import "github.com/aljrubior/anyctl/managers/entities"

type PrivateSpaceManager interface {
	GetPrivateSpace(ctx *entities.CurrentContextEntity, privateSpaceId string) (*entities.PrivateSpaceEntity, error)
	GetPrivateSpaces(ctx *entities.CurrentContextEntity) (*[]entities.PrivateSpaceEntity, error)
	FindPrivateSpaceByNameOrId(ctx *entities.CurrentContextEntity, privateSpaceId string) (*[]entities.PrivateSpaceEntity, error)

	GetFabrics(ctx *entities.CurrentContextEntity, privateSpaceId string) (*[]entities.FabricEntity, error)
	GetFabric(ctx *entities.CurrentContextEntity, privateSpaceId, fabricId string) (*entities.FabricEntity, error)
}
