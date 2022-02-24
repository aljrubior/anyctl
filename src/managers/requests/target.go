package requests

type Target struct {
	Provider           string             `json:"provider,omitempty"`
	TargetId           string             `json:"targetId,omitempty"`
	DeploymentSettings DeploymentSettings `json:"deploymentSettings,omitempty"`
	Replicas           int                `json:"replicas,omitempty"`
	Type               string             `json:"type,omitempty"`
}
