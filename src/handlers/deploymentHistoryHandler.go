package handlers

type DeploymentHistoryHandler interface {
	GetDeploymentHistory(deploymentName string) error
	Compare(deploymentName, compareWithSpecVersion string) error
}
