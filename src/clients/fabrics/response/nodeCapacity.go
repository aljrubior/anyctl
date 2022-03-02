package response

type NodeCapacity struct {
	Cpu       int    `json:"cpu"`
	CpuMillis int    `json:"cpuMillis"`
	Memory    string `json:"memory"`
	MemoryMi  int    `json:"memoryMi"`
	Pods      int    `json:"pods"`
}
