package response

type Environments struct {
	Type           string   `json:"type"`
	BusinessGroups []string `json:"businessGroups"`
}
