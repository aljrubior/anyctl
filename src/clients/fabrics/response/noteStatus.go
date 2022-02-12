package response

type NodeStatus struct {
	IsHealthy     bool `yaml:"isHealthy",json:"isHealthy"`
	IsReady       bool `yaml:"isReady",json:"isReady"`
	IsSchedulable bool `yaml:"isSchedulable",json:"isSchedulable"`
}
