package response

type FabricNode struct {
	Uid                      string       `yaml:"uid",json:"uid"`
	Name                     string       `yaml:"name",json:"name"`
	KubeletVersion           string       `yaml:"kubeletVersion",json:"kubeletVersion"`
	DockerVersion            string       `yaml:"dockerVersion",json:"dockerVersion"`
	Role                     string       `yaml:"role",json:"role"`
	Status                   NodeStatus   `yaml:"status",json:"status"`
	Capacity                 NodeCapacity `yaml:"capacity",json:"capacity"`
	AllocatedRequestCapacity NodeCapacity `yaml:"allocatedRequestCapacity",json:"allocatedRequestCapacity"`
	AllocatedLimitCapacity   NodeCapacity `yaml:"allocatedLimitCapacity",json:"allocatedLimitCapacity"`
}
