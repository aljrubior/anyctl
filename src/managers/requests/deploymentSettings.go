package requests

type DeploymentSettings struct {
	Resources                           *Resources `json:"resources,omitempty"`
	Clustered                           bool       `json:"clustered,omitempty"`
	EnforceDeployingReplicasAcrossNodes bool       `json:"enforceDeployingReplicasAcrossNodes,omitempty"`
	Http                                *Http      `json:"http,omitempty"`
	Jvm                                 *Jvm       `json:"jvm,omitempty"`
	RuntimeVersion                      string     `json:"runtimeVersion,omitempty"`
	LastMileSecurity                    bool       `json:"lastMileSecurity,omitempty"`
	ForwardSslSession                   bool       `json:"forwardSslSession,omitempty"`
	UpdateStrategy                      string     `json:"updateStrategy,omitempty"`
}
