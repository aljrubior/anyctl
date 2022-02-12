package response

type NodeCapacity struct {
	Cpu       int    `yaml:"cpu",json:"cpu"`
	CpuMillis int    `yaml:"cpuMillis",json:"cpuMillis"`
	Memory    string `yaml:"memory",json:"memory"`
	MemoryMi  int    `yaml:"memoryMi",json:"memoryMi"`
	Pods      int    `yaml:"pods",json:"pods"`
}
