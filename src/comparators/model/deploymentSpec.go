package model

type DeploymentSpec struct {
	Name           string      `json:"name"`
	Target         Target      `json:"target"`
	Application    Application `json:"application"`
	DesiredVersion string      `json:"desiredVersion"`
}
