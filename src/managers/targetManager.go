package managers

import (
	"github.com/aljrubior/anyctl/managers/entities"
)

type TargetManager interface {
	GetTargets(ctx *entities.CurrentContextEntity) (*[]entities.TargetEntity, error)
	FindTargetByName(ctx *entities.CurrentContextEntity, targetName string) (*entities.TargetEntity, *[]entities.TargetEntity, error)

	FindTargetsContainsName(ctx *entities.CurrentContextEntity, targetName string) (*[]entities.TargetEntity, error)
}
