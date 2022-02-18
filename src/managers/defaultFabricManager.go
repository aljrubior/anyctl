package managers

import (
	"github.com/aljrubior/anyctl/managers/entities"
	"github.com/aljrubior/anyctl/services"
)

func NewDefaultFabricManager(fabricService services.FabricService) DefaultFabricManager {
	return DefaultFabricManager{
		fabricService,
	}
}

type DefaultFabricManager struct {
	fabricService services.FabricService
}

func (this DefaultFabricManager) GetFabrics(ctx *entities.CurrentContextEntity) (*[]entities.FabricEntity, error) {

	resp, err := this.fabricService.GetFabrics(ctx.AuthorizationToken)

	if err != nil {
		return nil, err
	}

	return entities.NewFabricEntitiesBuilder(resp).Build(), nil
}

func (this DefaultFabricManager) GetFabric(ctx *entities.CurrentContextEntity, fabricId string) (*entities.FabricEntity, error) {

	resp, err := this.fabricService.GetFabric(ctx.AuthorizationToken, fabricId)

	if err != nil {
		return nil, err
	}

	return entities.NewFabricEntityBuilder(resp).Build(), nil
}

func (this DefaultFabricManager) FindFabricByNameOrId(ctx *entities.CurrentContextEntity, fabricId string) (*[]entities.FabricEntity, error) {

	resp, err := this.fabricService.GetFabricsByNameOrId(ctx.AuthorizationToken, fabricId)

	if err != nil {
		return nil, err
	}

	return entities.NewFabricEntitiesBuilder(resp).Build(), nil
}
