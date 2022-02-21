package response

type Application struct {
	Status        string                   `yaml:"status",json:"status"`
	DesiredState  string                   `yaml:"desiredState,omitempty",json:"desiredState,omitempty"`
	Ref           Asset                    `json:"ref"`
	Configuration ApplicationConfiguration `json:"configuration"`
}
