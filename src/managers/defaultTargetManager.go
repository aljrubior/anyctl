package managers

import (
	"github.com/aljrubior/anyctl/managers/entities"
	"github.com/aljrubior/anyctl/services"
	"strings"
)

func NewDefaultTargetManager(targetService services.TargetService) *DefaultTargetManager {
	return &DefaultTargetManager{
		targetService: targetService,
	}
}

type DefaultTargetManager struct {
	targetService services.TargetService
}

func (this DefaultTargetManager) GetTargets(ctx *entities.CurrentContextEntity) (*[]entities.TargetEntity, error) {

	resp, err := this.targetService.GetTargets(ctx.OrganizationId, ctx.EnvironmentId, ctx.AuthorizationToken)

	if err != nil {
		return nil, err
	}

	return entities.NewTargetEntitiesBuilder(resp).Build(), nil
}

func (this DefaultTargetManager) FindTargetByName(ctx *entities.CurrentContextEntity, targetName string) (*entities.TargetEntity, *[]entities.TargetEntity, error) {

	targets, err := this.GetTargets(ctx)

	if err != nil {
		return nil, nil, err
	}

	for _, v := range *targets {
		if v.GetName() == targetName {
			return &v, nil, nil
		}
	}

	return nil, targets, nil
}

func (this DefaultTargetManager) FindTargetsContainsName(ctx *entities.CurrentContextEntity, targetName string) (*[]entities.TargetEntity, error) {
	targets, err := this.GetTargets(ctx)

	if err != nil {
		return nil, err
	}

	var targetsFound []entities.TargetEntity

	for _, v := range *targets {
		if strings.Contains(v.GetName(), targetName) {
			targetsFound = append(targetsFound, v)
		}
	}

	return &targetsFound, nil
}

func (this *DefaultTargetManager) find(targets *[]entities.TargetEntity, targetName string) *entities.TargetEntity {

	for _, v := range *targets {
		if v.GetName() == targetName {
			return &v
		}
	}

	return nil
}
