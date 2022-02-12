package requests

type ResourceItem struct {
	Reserved string `json:"reserved,omitempty"`
	Limit    string `json:"limit,omitempty"`
}
