package errors

import (
	"fmt"
	"github.com/aljrubior/anyctl/managers/entities"
)

func NewSharedSpaceNotFoundError(targetName string, options *[]entities.SharedSpaceEntity) *SharedSpaceNotFoundError {
	return &SharedSpaceNotFoundError{
		targetName,
		options,
	}
}

type SharedSpaceNotFoundError struct {
	PrivateSpaceName string
	Options          *[]entities.SharedSpaceEntity
}

func (this *SharedSpaceNotFoundError) Error() string {
	return fmt.Sprintf("Shared Space '%s' not found", this.PrivateSpaceName)
}
