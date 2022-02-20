package conf

type DeploymentClientConfig struct {
	Protocol             string
	Host                 string
	Port                 int
	DeploymentsPath      string
	DeploymentPath       string
	UpdateDeploymentPath string
	SpecsPath            string
}
