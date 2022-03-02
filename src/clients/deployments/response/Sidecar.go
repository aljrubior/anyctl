package response

type Sidecar struct {
	Image     string    `json:"image"`
	Resources Resources `json:"resources"`
}
