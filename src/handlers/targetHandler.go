package handlers

import "github.com/aljrubior/anyctl/managers/entities"

type TargetHandler interface {
	GetTargets() error
	FindTargetsContainsName(targetName string) error
	GetSupportedRuntimes(targetName string) error
	GetDetails(targetName string) error
	GetAddresses(targetName string) error
	DescribeTarget(targetName string) error
	ThrowTargetNotFoundError(targetName string, targets *[]entities.TargetEntity) error
}
