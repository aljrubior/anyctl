package response

type Application struct {
	Status       string `yaml:"status",json:"status"`
	DesiredState string `yaml:"desiredState,omitempty",json:"desiredState,omitempty"`
	Asset        Asset  `yaml:"ref",json:"ref,omitempty"`
}
