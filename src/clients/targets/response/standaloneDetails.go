package response

type StandaloneDetails struct {
	RuntimeVersion string      `json:"runtimeVersion"`
	Type           string      `json:"type"`
	AgentVersion   string      `json:"agentVersion"`
	Addresses      []Addresses `json:"addresses"`
}
