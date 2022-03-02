package errors

import (
	"fmt"
	"github.com/aljrubior/anyctl/managers/entities"
)

func NewSchedulerNotFoundError(flowName string, options *[]entities.SchedulerEntity) *SchedulerNotFoundError {
	return &SchedulerNotFoundError{
		FlowName: flowName,
		Options:  options,
	}
}

type SchedulerNotFoundError struct {
	FlowName string
	Options  *[]entities.SchedulerEntity
}

func (this *SchedulerNotFoundError) Error() string {
	return fmt.Sprintf("Scheduler '%s' not found", this.FlowName)
}
