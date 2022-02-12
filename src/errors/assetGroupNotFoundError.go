package errors

import (
	"fmt"
	"github.com/aljrubior/anyctl/managers/entities"
)

func NewAssetGroupNotFoundError(groupId, assetName string, options *[]entities.AssetEntity) *AssetGroupNotFoundError {
	return &AssetGroupNotFoundError{
		groupId,
		assetName,
		options,
	}
}

type AssetGroupNotFoundError struct {
	GroupId   string
	AssetName string
	Options   *[]entities.AssetEntity
}

func (this *AssetGroupNotFoundError) Error() string {
	return fmt.Sprintf("Group '%s' not found for artifact '%s'", this.GroupId, this.AssetName)
}
