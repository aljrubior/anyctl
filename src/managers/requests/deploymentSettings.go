package requests

type DeploymentSettings struct {
	Jvm                                 Jvm       `json:"jvm,omitempty"`
	AnypointMonitoringScope             string    `json:"anypointMonitoringScope,omitempty"`
	Sidecars                            *Sidecars `json:"sidecars,omitempty"`
	UpdateStrategy                      string    `json:"updateStrategy,omitempty"`
	RuntimeVersion                      string    `json:"runtimeVersion,omitempty"`
	Clustered                           bool      `json:"clustered,omitempty"`
	ForwardSslSession                   bool      `json:"forwardSslSession,omitempty"`
	Http                                Http      `json:"http,omitempty"`
	Resources                           Resources `json:"resources,omitempty"`
	LastMileSecurity                    bool      `json:"lastMileSecurity,omitempty"`
	EnforceDeployingReplicasAcrossNodes bool      `json:"enforceDeployingReplicasAcrossNodes,omitempty"`
}
