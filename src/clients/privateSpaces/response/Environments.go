package response

type Environments struct {
	Type           string   `json:"type",yaml:"type"`
	BusinessGroups []string `json:"businessGroups",yaml:"businessGroups"`
}
