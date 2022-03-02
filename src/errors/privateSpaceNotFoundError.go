package errors

import (
	"fmt"
	"github.com/aljrubior/anyctl/managers/entities"
)

func NewPrivateSpaceNotFoundError(targetName string, options *[]entities.PrivateSpaceEntity) *PrivateSpaceNotFoundError {
	return &PrivateSpaceNotFoundError{
		targetName,
		options,
	}
}

type PrivateSpaceNotFoundError struct {
	PrivateSpaceName string
	Options          *[]entities.PrivateSpaceEntity
}

func (this *PrivateSpaceNotFoundError) Error() string {
	return fmt.Sprintf("Private Space '%s' not found", this.PrivateSpaceName)
}
