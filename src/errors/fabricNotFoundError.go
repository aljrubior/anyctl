package errors

import (
	"fmt"
	"github.com/aljrubior/anyctl/managers/entities"
)

func NewFabricNotFoundError(fabricId string, options *[]entities.FabricEntity) *FabricNotFoundError {
	return &FabricNotFoundError{
		fabricId,
		options,
	}
}

type FabricNotFoundError struct {
	FabricId string
	Options  *[]entities.FabricEntity
}

func (this *FabricNotFoundError) Error() string {
	return fmt.Sprintf("Fabric '%s' not found", this.FabricId)
}
