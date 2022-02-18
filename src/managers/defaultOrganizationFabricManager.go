package managers

import (
	errors2 "github.com/aljrubior/anyctl/errors"
	"github.com/aljrubior/anyctl/managers/entities"
	"github.com/aljrubior/anyctl/managers/wrappers"
	"github.com/aljrubior/anyctl/services"
	"strings"
)

func NewDefaultOrganizationRuntimeFabricManager(runtimeFabricService services.OrganizationRuntimeFabricService) DefaultOrganizationRuntimeFabricManager {
	return DefaultOrganizationRuntimeFabricManager{
		runtimeFabricService,
	}
}

type DefaultOrganizationRuntimeFabricManager struct {
	runtimeFabricService services.OrganizationRuntimeFabricService
}

func (this DefaultOrganizationRuntimeFabricManager) GetFabrics(ctx *entities.CurrentContextEntity) (*[]entities.OrganizationFabricEntity, error) {

	resp, err := this.runtimeFabricService.GetFabrics(ctx.OrganizationId, ctx.EnvironmentId, ctx.AuthorizationToken)

	if err != nil {
		return nil, err
	}

	return entities.NewOrganizationFabricEntitiesBuilder(resp).Build(), nil
}

func (this DefaultOrganizationRuntimeFabricManager) GetFabric(ctx *entities.CurrentContextEntity, targetId string) (*entities.OrganizationFabricEntity, error) {

	resp, err := this.runtimeFabricService.GetFabric(ctx.OrganizationId, ctx.EnvironmentId, ctx.AuthorizationToken, targetId)

	if err != nil {
		return nil, err
	}

	return entities.NewOrganizationFabricEntityBuilder(resp).Build(), nil
}

func (this DefaultOrganizationRuntimeFabricManager) GetFabricTarget(ctx *entities.CurrentContextEntity, targetId string) (*entities.FabricTargetEntity, error) {

	resp, err := this.runtimeFabricService.GetTarget(ctx.OrganizationId, ctx.EnvironmentId, ctx.AuthorizationToken, targetId)

	if err != nil {
		return nil, err
	}

	return entities.NewFabricTargetEntityBuilder(resp).Build(), nil
}

func (this DefaultOrganizationRuntimeFabricManager) GetFabricTargets(ctx *entities.CurrentContextEntity) (*[]entities.FabricTargetEntity, error) {

	resp, err := this.runtimeFabricService.GetTargets(ctx.OrganizationId, ctx.EnvironmentId, ctx.AuthorizationToken)

	if err != nil {
		return nil, err
	}

	return entities.NewFabricTargetEntitiesBuilder(resp).Build(), nil
}

func (this DefaultOrganizationRuntimeFabricManager) FindFabricByName(ctx *entities.CurrentContextEntity, fabricName string) (*entities.OrganizationFabricEntity, *[]entities.OrganizationFabricEntity, error) {

	fabrics, err := this.GetFabrics(ctx)

	if err != nil {
		return nil, nil, err
	}

	for _, v := range *fabrics {
		if v.Name == fabricName {
			return &v, nil, nil
		}
	}

	return nil, fabrics, nil
}

func (this DefaultOrganizationRuntimeFabricManager) FindFabricTargetByName(ctx *entities.CurrentContextEntity, targetName string) (*entities.FabricTargetEntity, *[]entities.FabricTargetEntity, error) {

	targets, err := this.GetFabricTargets(ctx)

	if err != nil {
		return nil, nil, err
	}

	for _, v := range *targets {
		if v.Name == targetName {
			return &v, nil, nil
		}
	}

	return nil, targets, nil
}

func (this DefaultOrganizationRuntimeFabricManager) FindExactMatch(ctx *entities.CurrentContextEntity, targetName string) (*entities.FabricTargetEntity, error) {

	resp, err := this.GetFabricTargets(ctx)

	if err != nil {
		return nil, err
	}

	targetWrapper := wrappers.NewFabricTargetEntitiesWrapper(resp)

	target := targetWrapper.GetTargetByName(targetName)

	if target == nil {
		return nil, nil
	}

	return target, nil
}

func (this DefaultOrganizationRuntimeFabricManager) FindRuntimeFabricContainsName(ctx *entities.CurrentContextEntity, runtimeFabricName string) (*[]entities.OrganizationFabricEntity, error) {
	targets, err := this.GetFabrics(ctx)

	if err != nil {
		return nil, err
	}

	var runtimeFabricsFound []entities.OrganizationFabricEntity

	for _, v := range *targets {
		if strings.Contains(v.Name, runtimeFabricName) {
			runtimeFabricsFound = append(runtimeFabricsFound, v)
		}
	}

	return &runtimeFabricsFound, nil
}

func (this DefaultOrganizationRuntimeFabricManager) FindRuntimeLatestSupportedVersion(ctx *entities.CurrentContextEntity, targetName string) (*string, error) {

	resp, err := this.GetFabricTargets(ctx)

	if err != nil {
		return nil, err
	}

	targetWrapper := wrappers.NewFabricTargetEntitiesWrapper(resp)

	if !targetWrapper.ExistsTarget(targetName) {
		return nil, this.ThrowFabricTargetNotFoundError(targetName, resp)
	}

	versionRef := targetWrapper.GetLatestRuntimeVersion(targetName)

	if versionRef == "" {
		return nil, this.ThrowRuntimeVersionNotFoundError(targetName)
	}

	return &versionRef, nil
}

func (this DefaultOrganizationRuntimeFabricManager) ThrowFabricTargetNotFoundError(targetName string, options *[]entities.FabricTargetEntity) error {
	return errors2.NewFabricTargetNotFoundError(targetName, options)
}

func (this DefaultOrganizationRuntimeFabricManager) ThrowRuntimeVersionNotFoundError(targetName string) error {
	return errors2.NewRuntimeVersionNotFoundError(targetName)
}
