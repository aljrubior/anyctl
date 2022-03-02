package model

type Target struct {
	Provider           string             `json:"provider"`
	TargetId           string             `json:"targetId"`
	DeploymentSettings DeploymentSettings `json:"deploymentSettings"`
	Replicas           int                `json:"replicas"`
	Type               string             `json:"type"`
}
