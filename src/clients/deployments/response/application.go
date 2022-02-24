package response

type Application struct {
	Ref           Asset                    `json:"ref"`
	Status        string                   `json:"status"`
	DesiredState  string                   `json:"desiredState,omitempty"`
	Configuration ApplicationConfiguration `json:"configuration"`
}
