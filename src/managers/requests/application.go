package requests

type Application struct {
	Ref           *ArtifactRef              `json:"ref,omitempty"`
	Assets        []interface{}             `json:"assets,omitempty"`
	DesiredState  string                    `json:"desiredState,omitempty"`
	Configuration *ApplicationConfiguration `json:"configuration,omitempty"`
}
