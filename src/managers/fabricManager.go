package managers

import "github.com/aljrubior/anyctl/managers/entities"

type FabricManager interface {
	GetFabric(ctx *entities.CurrentContextEntity, fabricId string) (*entities.FabricEntity, error)
	GetFabrics(ctx *entities.CurrentContextEntity) (*[]entities.FabricEntity, error)
	FindFabricByNameOrId(ctx *entities.CurrentContextEntity, fabricId string) (*[]entities.FabricEntity, error)
}
