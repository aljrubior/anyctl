package response

type MuleApplicationDefault struct {
	Sidecars                Sidecars `yaml:"sidecars",json:"sidecars"`
	AnypointMonitoringScope string   `yaml:"anypointMonitoringScope",json:"anypointMonitoringScope"`
}
