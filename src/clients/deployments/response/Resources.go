package response

type Resources struct {
	Cpu    ResourceItem `json:"cpu"`
	Memory ResourceItem `json:"memory"`
}
