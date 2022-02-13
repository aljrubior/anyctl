package managers

import "github.com/aljrubior/anyctl/managers/entities"

type OrganizationRuntimeFabricManager interface {
	GetFabrics(ctx *entities.CurrentContextEntity) (*[]entities.OrganizationFabricEntity, error)
	GetFabric(ctx *entities.CurrentContextEntity, targetId string) (*entities.OrganizationFabricEntity, error)
	GetFabricTarget(ctx *entities.CurrentContextEntity, targetId string) (*entities.FabricTargetEntity, error)
	GetFabricTargets(ctx *entities.CurrentContextEntity) (*[]entities.FabricTargetEntity, error)
	FindFabricByName(ctx *entities.CurrentContextEntity, fabricName string) (*entities.OrganizationFabricEntity, *[]entities.OrganizationFabricEntity, error)
	FindFabricTargetByName(ctx *entities.CurrentContextEntity, targetName string) (*entities.FabricTargetEntity, *[]entities.FabricTargetEntity, error)
	FindExactMatch(ctx *entities.CurrentContextEntity, targetName string) (*entities.FabricTargetEntity, error)
	FindRuntimeLatestSupportedVersion(ctx *entities.CurrentContextEntity, targetName string) (*string, error)

	FindRuntimeFabricContainsName(ctx *entities.CurrentContextEntity, runtimeFabricName string) (*[]entities.OrganizationFabricEntity, error)
}
