package services

import (
	"github.com/aljrubior/anyctl/clients/sharedspaces/response"
)

type SharedSpaceService interface {
	GetSharedSpaces(token string) (*[]response.SharedSpaceResponse, error)
	GetSharedSpace(token, targetId string) (*response.SharedSpaceResponse, error)
}
