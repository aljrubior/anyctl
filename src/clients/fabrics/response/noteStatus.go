package response

type NodeStatus struct {
	IsHealthy     bool `json:"isHealthy"`
	IsReady       bool `json:"isReady"`
	IsSchedulable bool `json:"isSchedulable"`
}
