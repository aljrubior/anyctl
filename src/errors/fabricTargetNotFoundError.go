package errors

import (
	"fmt"
	"github.com/aljrubior/anyctl/managers/entities"
)

func NewFabricTargetNotFoundError(targetName string, options *[]entities.FabricTargetEntity) *FabricTargetNotFoundError {
	return &FabricTargetNotFoundError{
		targetName,
		options,
	}
}

type FabricTargetNotFoundError struct {
	TargetName string
	Options    *[]entities.FabricTargetEntity
}

func (this *FabricTargetNotFoundError) Error() string {
	return fmt.Sprintf("Target '%s' not found", this.TargetName)
}
