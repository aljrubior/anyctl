package errors

import (
	"fmt"
	"github.com/aljrubior/anyctl/managers/entities"
)

func NewRtfTargetNotFoundError(targetName string, options *[]entities.RtfTargetEntity) *RtfTargetNotFoundError {
	return &RtfTargetNotFoundError{
		targetName,
		options,
	}
}

type RtfTargetNotFoundError struct {
	TargetName string
	Options    *[]entities.RtfTargetEntity
}

func (this *RtfTargetNotFoundError) Error() string {
	return fmt.Sprintf("Rtf Target '%s' not found", this.TargetName)
}
