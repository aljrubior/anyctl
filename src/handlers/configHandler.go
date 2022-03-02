package handlers

type ConfigHandler interface {
	Initialize(username, password string) error
	RefreshAccessToken() error
	SetCurrentEnvironment(name string) error
	PrintCurrentContext() error
}
