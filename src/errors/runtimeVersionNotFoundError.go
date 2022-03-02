package errors

import (
	"fmt"
)

func NewRuntimeVersionNotFoundError(targetName string) *RuntimeVersionNotFoundError {
	return &RuntimeVersionNotFoundError{
		targetName,
	}
}

type RuntimeVersionNotFoundError struct {
	TargetName string
}

func (this *RuntimeVersionNotFoundError) Error() string {
	return fmt.Sprintf("Runtime version for target '%s' not found", this.TargetName)
}
