package managers

import (
	"github.com/aljrubior/anyctl/managers/entities"
	"github.com/aljrubior/anyctl/services"
	"strings"
)

func NewDefaultSharedSpaceManager(sharedSpaceService services.SharedSpaceService) DefaultSharedSpaceManager {
	return DefaultSharedSpaceManager{
		sharedSpaceService,
	}
}

type DefaultSharedSpaceManager struct {
	sharedSpaceService services.SharedSpaceService
}

func (this DefaultSharedSpaceManager) GetSharedSpace(ctx *entities.CurrentContextEntity, sharedSpaceId string) (*entities.SharedSpaceEntity, error) {

	resp, err := this.sharedSpaceService.GetSharedSpace(ctx.AuthorizationToken, sharedSpaceId)

	if err != nil {
		return nil, err
	}

	return entities.NewSharedSpaceEntityBuilder(resp).Build(), nil
}

func (this DefaultSharedSpaceManager) GetSharedSpaces(ctx *entities.CurrentContextEntity) (*[]entities.SharedSpaceEntity, error) {

	resp, err := this.sharedSpaceService.GetSharedSpaces(ctx.AuthorizationToken)

	if err != nil {
		return nil, err
	}

	return entities.NewSharedSpaceEntitiesBuilder(resp).Build(), nil
}

func (this DefaultSharedSpaceManager) FindSharedSpaceByName(ctx *entities.CurrentContextEntity, ssName string) (*entities.SharedSpaceEntity, *[]entities.SharedSpaceEntity, error) {

	sharedSpaces, err := this.GetSharedSpaces(ctx)

	if err != nil {
		return nil, nil, err
	}

	for _, v := range *sharedSpaces {
		if v.Name == ssName {
			return &v, nil, nil
		}
	}

	return nil, sharedSpaces, nil
}

func (this DefaultSharedSpaceManager) FindSharedSpaceContainsName(ctx *entities.CurrentContextEntity, sharedSpaceName string) (*[]entities.SharedSpaceEntity, error) {

	sharedSpaces, err := this.GetSharedSpaces(ctx)

	if err != nil {
		return nil, err
	}

	var sharedSpacesFound []entities.SharedSpaceEntity

	for _, v := range *sharedSpaces {
		if strings.Contains(v.Name, sharedSpaceName) {
			sharedSpacesFound = append(sharedSpacesFound, v)
		}
	}

	return &sharedSpacesFound, nil
}
