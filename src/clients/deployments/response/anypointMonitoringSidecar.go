package response

type AnypointMonitoringSidecar struct {
	Image     string    `json:"image"`
	Resources Resources `json:"resources"`
}
