package errors

import (
	"fmt"
	"github.com/aljrubior/anyctl/managers/entities"
)

func NewAssetNotFoundError(assetName string, options *[]entities.AssetEntity) *AssetNotFoundError {
	return &AssetNotFoundError{
		AssetName: assetName,
		Options:   options,
	}
}

type AssetNotFoundError struct {
	AssetName string
	Reason    string
	Options   *[]entities.AssetEntity
}

func (this *AssetNotFoundError) WithReason(reason string) *AssetNotFoundError {
	this.Reason = reason
	return this
}

func (this *AssetNotFoundError) Error() string {

	if this.Reason == "" {
		return fmt.Sprintf("Asset '%s' not found", this.AssetName)
	}

	return fmt.Sprintf("Asset '%s' not found. Reason: %s", this.AssetName, this.Reason)
}
