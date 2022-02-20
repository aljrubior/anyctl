package response

type DeploymentSpecResponse struct {
	CreatedAt   int64       `json:"createdAt"`
	Target      interface{} `json:"target"`
	Application interface{} `json:"application"`
	Version     string      `json:"version"`
}
