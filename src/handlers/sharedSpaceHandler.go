package handlers

import "github.com/aljrubior/anyctl/managers/entities"

type SharedSpaceHandler interface {
	GetSharedSpaces() error
	GetSharedSpace(ssName string) error
	DescribeSharedSpace(sharedSpaceName string) error

	ThrowNewSharedSpaceNotFoundError(ssName string, options *[]entities.SharedSpaceEntity) error
}
