package response

type Resources struct {
	Cpu    ResourceItem `yaml:"cpu",json:"cpu"`
	Memory ResourceItem `yaml:"memory",json:"memory"`
}
