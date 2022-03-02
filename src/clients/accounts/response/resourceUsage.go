package response

type ResourceUsage struct {
	Assigned   float64 `json:"assigned"`
	Reassigned float64 `json:"reassigned"`
}
