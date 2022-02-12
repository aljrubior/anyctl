package errors

import (
	"fmt"
	"github.com/aljrubior/anyctl/managers/entities"
)

func NewTargetNotFoundError(targetName string, options *[]entities.TargetEntity) *TargetNotFoundError {
	return &TargetNotFoundError{
		targetName,
		options,
	}
}

type TargetNotFoundError struct {
	TargetName string
	Options    *[]entities.TargetEntity
}

func (this *TargetNotFoundError) Error() string {
	return fmt.Sprintf("Target '%s' not found", this.TargetName)
}
