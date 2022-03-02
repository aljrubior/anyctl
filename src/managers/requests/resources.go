package requests

type Resources struct {
	Cpu    ResourceItem `json:"cpu,omitempty"`
	Memory ResourceItem `json:"memory,omitempty"`
}
