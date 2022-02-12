package response

type Replica struct {
	Id                       string `yaml:"id",json:"id"`
	State                    string `yaml:"state",json:"state"`
	DeploymentLocation       string `yaml:"deploymentLocation",json:"deploymentLocation"`
	CurrentDeploymentVersion string `yaml:"currentDeploymentVersion",json:"currentDeploymentVersion"`
	Reason                   string `yaml:"reason",json:"reason"`
}
