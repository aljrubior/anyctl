package conf

func NewAppConfig() *AppConfig {
	return &AppConfig{}
}

type AppConfig struct {
	configDir          string
	anypointConfigFile string
	anyctlConfigFile   string
}

func (this AppConfig) ConfigDir() string {
	return ".anypoint"
}

func (this AppConfig) AnypointConfigFile() string {
	return "anyconfig"
}

func (this AppConfig) AnyctlConfigFile() string {
	return "anyctl.yaml"
}
