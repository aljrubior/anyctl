package response

type Target struct {
	Provider           string             `yaml:"provider",json:"provider"`
	TargetId           string             `yaml:"targetId",json:"targetId"`
	DeploymentSettings DeploymentSettings `yaml:"deploymentSettings",json:"deploymentSettings"`
	Replicas           int                `yaml:"replicas",json:"replicas"`
	Type               string             `yaml:"type",json:"type"`
}
