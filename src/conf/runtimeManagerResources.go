package conf

type RuntimeManagerResources struct {
	Deployments    DeploymentResource     `yaml:"deployments"`
	DeploymentLogs DeploymentLogsResource `yaml:"deploymentLogs"`
	Schedulers     SchedulerResource      `yaml:"schedulers"`
	Assets         AssetResource          `yaml:"assets"`
	Targets        TargetResource         `yaml:"targets"`
	RuntimeFabrics RuntimeFabricResource  `yaml:"runtimeFabrics"`
}
