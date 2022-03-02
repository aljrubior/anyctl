package entities

import (
	"github.com/aljrubior/anyctl/clients/accounts/response"
)

func NewLoginEntityBuilder(token string, profile *response.Profile) *LoginEntityBuilder {
	return &LoginEntityBuilder{
		token:   token,
		profile: profile,
	}
}

type LoginEntityBuilder struct {
	token   string
	profile *response.Profile
}

func (this LoginEntityBuilder) Build() *LoginEntity {
	return &LoginEntity{
		OrganizationProfile: *this.profile,
		Token:               this.token,
	}
}
