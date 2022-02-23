package response

type Application struct {
	Status        string                   `json:"status"`
	DesiredState  string                   `json:"desiredState,omitempty"`
	Ref           Asset                    `json:"ref"`
	Configuration ApplicationConfiguration `json:"configuration"`
}
