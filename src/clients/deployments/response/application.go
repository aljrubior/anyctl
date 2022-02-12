package response

type Application struct {
	Status       string `json:"status"`
	DesiredState string `yaml:"desiredState,omitempty",json:"desiredState,omitempty"`
	Asset        Asset  `json:"ref"`
}
