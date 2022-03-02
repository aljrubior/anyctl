package requests

func NewLoginRequestBuilder(username string, password string) *LoginRequestBuilder {
	return &LoginRequestBuilder{
		username,
		password,
	}
}

type LoginRequestBuilder struct {
	username string
	password string
}

func (this *LoginRequestBuilder) Build() *LoginRequest {

	return &LoginRequest{
		Username: this.username,
		Password: this.password,
	}
}
