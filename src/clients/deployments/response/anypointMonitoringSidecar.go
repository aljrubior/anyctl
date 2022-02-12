package response

type AnypointMonitoringSidecar struct {
	Image     string    `yaml:"image",json:"image"`
	Resources Resources `yaml:"resources",json:"resources"`
}
