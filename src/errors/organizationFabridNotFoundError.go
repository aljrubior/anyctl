package errors

import (
	"fmt"
	"github.com/aljrubior/anyctl/managers/entities"
)

func NewOrganizationFabricNotFoundError(runtimeFabricName string, options *[]entities.OrganizationFabricEntity) *OrganizationFabricNotFoundError {
	return &OrganizationFabricNotFoundError{
		runtimeFabricName,
		options,
	}
}

type OrganizationFabricNotFoundError struct {
	RuntimeFabricName string
	Options           *[]entities.OrganizationFabricEntity
}

func (this *OrganizationFabricNotFoundError) Error() string {
	return fmt.Sprintf("Runtime Fabric '%s' not found", this.RuntimeFabricName)
}
