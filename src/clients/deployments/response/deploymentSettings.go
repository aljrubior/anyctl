package response

type DeploymentSettings struct {
	Jvm                                 Jvm       `yaml:"jvm,omitempty",json:"jvm,omitempty"`
	AnypointMonitoringScope             string    `yaml:"anypointMonitoringScope,omitempty",json:"anypointMonitoringScope,omitempty"`
	Sidecars                            Sidecars  `yaml:"sidecars,omitempty",json:"sidecars,omitempty"`
	UpdateStrategy                      string    `yaml:"updateStrategy,omitempty",json:"updateStrategy,omitempty"`
	RuntimeVersion                      string    `yaml:"runtimeVersion,omitempty",json:"runtimeVersion,omitempty"`
	Clustered                           bool      `yaml:"clustered,omitempty",json:"clustered,omitempty"`
	ForwardSslSession                   bool      `yaml:"forwardSslSession,omitempty",json:"forwardSslSession,omitempty"`
	Http                                Http      `yaml:"http,omitempty",json:"http,omitempty"`
	Resources                           Resources `yaml:"resources,omitempty",json:"resources,omitempty"`
	LastMileSecurity                    bool      `yaml:"lastMileSecurity,omitempty",json:"lastMileSecurity,omitempty"`
	EnforceDeployingReplicasAcrossNodes bool      `yaml:"enforceDeployingReplicasAcrossNodes,omitempty",json:"enforceDeployingReplicasAcrossNodes,omitempty"`
}
