package response

type Runtime struct {
	Type     string           `json:"type"`
	Versions []RuntimeVersion `json:"versions"`
}
