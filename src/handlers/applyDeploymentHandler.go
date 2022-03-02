package handlers

import "github.com/aljrubior/anyctl/managers/entities"

type ApplyDeploymentHandler interface {
	Apply(filePath string) error
	ShowApplyPlan(filePath string) error

	ThrowTargetNotFoundError(targetId string, options *[]entities.TargetEntity) error
}
