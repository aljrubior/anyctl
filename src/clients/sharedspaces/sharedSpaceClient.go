package sharedspaces

import (
	"github.com/aljrubior/anyctl/clients/sharedspaces/response"
)

type SharedSpaceClient interface {
	GetSharedSpaces(token string) (*[]response.SharedSpaceResponse, error)
	GetSharedSpace(token, sharedSpaceId string) (*response.SharedSpaceResponse, error)
}
