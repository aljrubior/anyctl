package managers

import (
	"github.com/aljrubior/anyctl/managers/entities"
	"github.com/aljrubior/anyctl/services"
)

func NewDefaultPrivateSpaceManager(privateSpaceService services.PrivateSpaceService) *DefaultPrivateSpaceManager {
	return &DefaultPrivateSpaceManager{
		privateSpaceService,
	}
}

type DefaultPrivateSpaceManager struct {
	privateSpaceService services.PrivateSpaceService
}

func (this DefaultPrivateSpaceManager) GetPrivateSpaces(ctx *entities.CurrentContextEntity) (*[]entities.PrivateSpaceEntity, error) {

	resp, err := this.privateSpaceService.GetPrivateSpaces(ctx.AuthorizationToken)

	if err != nil {
		return nil, err
	}

	return entities.NewPrivateSpaceEntitiesBuilder(resp).Build(), nil
}

func (this DefaultPrivateSpaceManager) GetPrivateSpace(ctx *entities.CurrentContextEntity, privateSpaceId string) (*entities.PrivateSpaceEntity, error) {

	resp, err := this.privateSpaceService.GetPrivateSpace(ctx.AuthorizationToken, privateSpaceId)

	if err != nil {
		return nil, err
	}

	return entities.NewPrivateSpaceEntityBuilder(resp).Build(), nil
}

func (this DefaultPrivateSpaceManager) FindPrivateSpaceByNameOrId(ctx *entities.CurrentContextEntity, privateSpaceId string) (*[]entities.PrivateSpaceEntity, error) {

	resp, err := this.privateSpaceService.GetPrivateSpacesByNameOrId(ctx.AuthorizationToken, privateSpaceId)

	if err != nil {
		return nil, err
	}

	return entities.NewPrivateSpaceEntitiesBuilder(resp).Build(), nil
}

func (this DefaultPrivateSpaceManager) GetFabrics(ctx *entities.CurrentContextEntity, privateSpaceId string) (*[]entities.FabricEntity, error) {

	resp, err := this.privateSpaceService.GetFabrics(ctx.AuthorizationToken, privateSpaceId)

	if err != nil {
		return nil, err
	}

	return entities.NewFabricEntitiesBuilder(resp).Build(), nil
}

func (this DefaultPrivateSpaceManager) GetFabric(ctx *entities.CurrentContextEntity, privateSpaceId, fabricId string) (*entities.FabricEntity, error) {

	resp, err := this.privateSpaceService.GetFabric(ctx.AuthorizationToken, privateSpaceId, fabricId)

	if err != nil {
		return nil, err
	}

	return entities.NewFabricEntityBuilder(resp).Build(), nil
}
