package conf

type DeploymentResource struct {
	DeploymentPath       string `yaml:"deploymentPath"`
	DeploymentsPath      string `yaml:"deploymentsPath"`
	UpdateDeploymentPath string `yaml:"updateDeploymentsPath"`
}
