package handlers

type LoginHandler interface {
	Login(username, password string) error
}
