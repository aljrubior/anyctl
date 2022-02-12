package conf

type RuntimeManagerResources struct {
	Deployments    DeploymentResource    `yaml:"deployments"`
	Schedulers     SchedulerResource     `yaml:"schedulers"`
	Assets         AssetResource         `yaml:"assets"`
	Targets        TargetResource        `yaml:"targets"`
	RuntimeFabrics RuntimeFabricResource `yaml:"runtimeFabrics"`
}
