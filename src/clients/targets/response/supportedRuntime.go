package response

type SupportedRuntime struct {
	BaseVersion string `json:"baseVersion"`
	Tag         string `json:"tag"`
	MinimumTag  string `json:"minimumTag"`
}
