package requests

type DeploymentRequest struct {
	Id          string      `json:"id,omitempty"`
	Name        string      `json:"name,omitempty"`
	Labels      []string    `json:"labels,omitempty"`
	Target      Target      `json:"target,omitempty"`
	Application Application `json:"application,omitempty"`
}
