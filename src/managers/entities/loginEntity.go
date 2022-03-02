package entities

import (
	"github.com/aljrubior/anyctl/clients/accounts/response"
)

type LoginEntity struct {
	Token               string
	OrganizationProfile response.Profile
}
