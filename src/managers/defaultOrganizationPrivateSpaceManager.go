package managers

import (
	"github.com/aljrubior/anyctl/managers/entities"
	"github.com/aljrubior/anyctl/services"
	"strings"
)

func NewDefaultOrganizationPrivateSpaceManager(privateSpaceService services.OrganizationPrivateSpaceService) *DefaultOrganizationPrivateSpaceManager {
	return &DefaultOrganizationPrivateSpaceManager{
		privateSpaceService,
	}
}

type DefaultOrganizationPrivateSpaceManager struct {
	privateSpaceService services.OrganizationPrivateSpaceService
}

func (this DefaultOrganizationPrivateSpaceManager) GetPrivateSpaces(ctx *entities.CurrentContextEntity) (*[]entities.OrganizationPrivateSpaceEntity, error) {

	resp, err := this.privateSpaceService.GetPrivateSpaces(ctx.OrganizationId, ctx.EnvironmentId, ctx.AuthorizationToken)

	if err != nil {
		return nil, err
	}

	return entities.NewOrganizationPrivateSpaceEntitiesBuilder(resp).Build(), nil
}

func (this DefaultOrganizationPrivateSpaceManager) GetFabrics(ctx *entities.CurrentContextEntity, privateSpaceId string) (*[]entities.OrganizationPrivateSpaceFabricEntity, error) {

	resp, err := this.privateSpaceService.GetFabrics(ctx.OrganizationId, ctx.EnvironmentId, ctx.AuthorizationToken, privateSpaceId)

	if err != nil {
		return nil, err
	}

	return entities.NewOrganizationPrivateSpaceFabricEntitiesBuilder(resp).Build(), nil
}

func (this DefaultOrganizationPrivateSpaceManager) FindPrivateSpaceByName(ctx *entities.CurrentContextEntity, psName string) (*entities.OrganizationPrivateSpaceEntity, *[]entities.OrganizationPrivateSpaceEntity, error) {

	privateSpaces, err := this.GetPrivateSpaces(ctx)

	if err != nil {
		return nil, nil, err
	}

	for _, v := range *privateSpaces {
		if v.Name == psName {
			return &v, nil, nil
		}
	}

	return nil, privateSpaces, nil
}

func (this DefaultOrganizationPrivateSpaceManager) FindPrivateSpaceContainsName(ctx *entities.CurrentContextEntity, privateSpaceName string) (*[]entities.OrganizationPrivateSpaceEntity, error) {

	privateSpaces, err := this.GetPrivateSpaces(ctx)

	if err != nil {
		return nil, err
	}

	var privateSpacesFound []entities.OrganizationPrivateSpaceEntity

	for _, v := range *privateSpaces {
		if strings.Contains(v.Name, privateSpaceName) {
			privateSpacesFound = append(privateSpacesFound, v)
		}
	}

	return &privateSpacesFound, nil
}
