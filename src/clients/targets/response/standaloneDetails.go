package response

type StandaloneDetails struct {
	RuntimeVersion string      `yaml:"runtimeVersion",json:"runtimeVersion"`
	Type           string      `yaml:"type",json:"type"`
	AgentVersion   string      `yaml:"agentVersion",json:"agentVersion"`
	Addresses      []Addresses `yaml:"addresses",json:"addresses"`
}
