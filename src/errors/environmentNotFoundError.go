package errors

import (
	"fmt"
	"github.com/aljrubior/anyctl/managers/entities"
)

func NewEnvironmentNotFoundError(environmentName string, organization *entities.OrganizationEntity) *EnvironmentNotFoundError {
	return &EnvironmentNotFoundError{
		environmentName,
		organization,
	}
}

type EnvironmentNotFoundError struct {
	EnvironmentName string
	Organization    *entities.OrganizationEntity
}

func (this *EnvironmentNotFoundError) Error() string {
	return fmt.Sprintf("Environment '%s' not found", this.EnvironmentName)
}
