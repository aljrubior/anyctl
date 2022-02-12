package managers

import (
	"github.com/aljrubior/anyctl/managers/entities"
)

type LoginManager interface {
	Login(username, password string) (*entities.LoginEntity, error)
	RefreshAccessToken(username, password string) (string, error)
}
