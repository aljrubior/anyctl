package model

type Application struct {
	DesiredState  string                   `json:"desiredState,omitempty"`
	Ref           Asset                    `json:"ref"`
	Configuration ApplicationConfiguration `json:"configuration"`
}
