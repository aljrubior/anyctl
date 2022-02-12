package errors

import (
	"fmt"
	"github.com/aljrubior/anyctl/managers/entities"
)

func NewOrganizationPrivateSpaceNotFoundError(targetName string, options *[]entities.OrganizationPrivateSpaceEntity) *OrganizationPrivateSpaceNotFoundError {
	return &OrganizationPrivateSpaceNotFoundError{
		targetName,
		options,
	}
}

type OrganizationPrivateSpaceNotFoundError struct {
	PrivateSpaceName string
	Options          *[]entities.OrganizationPrivateSpaceEntity
}

func (this *OrganizationPrivateSpaceNotFoundError) Error() string {
	return fmt.Sprintf("Private Space '%s' not found", this.PrivateSpaceName)
}
