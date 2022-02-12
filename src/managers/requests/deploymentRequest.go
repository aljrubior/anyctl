package requests

type DeploymentRequest struct {
	Name        string       `json:"name,omitempty"`
	Labels      []string     `json:"labels,omitempty"`
	Target      *Target      `json:"target,omitempty"`
	Application *Application `json:"application,omitempty"`
}
