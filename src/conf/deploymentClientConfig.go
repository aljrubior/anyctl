package conf

type DeploymentClientConfig struct {
	Protocol                     string
	Host                         string
	Port                         int
	DeploymentsPathTemplate      string
	DeploymentPathTemplate       string
	UpdateDeploymentPathTemplate string
}
