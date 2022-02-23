package response

type Replica struct {
	Id                       string `json:"id"`
	State                    string `json:"state"`
	DeploymentLocation       string `json:"deploymentLocation"`
	CurrentDeploymentVersion string `json:"currentDeploymentVersion"`
	Reason                   string `json:"reason"`
}
