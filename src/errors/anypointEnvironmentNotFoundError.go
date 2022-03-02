package errors

import (
	"fmt"
	"github.com/aljrubior/anyctl/model"
)

func NewAnypointEnvironmentNotFoundError(environmentName string, options *[]model.Environment) *AnypointEnvironmentNotFoundError {
	return &AnypointEnvironmentNotFoundError{
		environmentName,
		options,
	}
}

type AnypointEnvironmentNotFoundError struct {
	EnvironmentName string
	Options         *[]model.Environment
}

func (this *AnypointEnvironmentNotFoundError) Error() string {
	return fmt.Sprintf("Environment '%s' not found", this.EnvironmentName)
}
