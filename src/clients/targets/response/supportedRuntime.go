package response

type SupportedRuntime struct {
	BaseVersion string `yaml:"baseVersion",json:"baseVersion"`
	Tag         string `yaml:"tag",json:"tag"`
	MinimumTag  string `yaml:"minimumTag",json:"minimumTag"`
}
