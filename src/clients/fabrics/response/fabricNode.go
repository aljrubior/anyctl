package response

type FabricNode struct {
	Uid                      string       `json:"uid"`
	Name                     string       `json:"name"`
	KubeletVersion           string       `json:"kubeletVersion"`
	DockerVersion            string       `json:"dockerVersion"`
	Role                     string       `json:"role"`
	Status                   NodeStatus   `json:"status"`
	Capacity                 NodeCapacity `json:"capacity"`
	AllocatedRequestCapacity NodeCapacity `json:"allocatedRequestCapacity"`
	AllocatedLimitCapacity   NodeCapacity `json:"allocatedLimitCapacity"`
}
