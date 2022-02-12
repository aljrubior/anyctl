package response

type Http struct {
	Inbound Inbound `yaml:"inbound,omitempty",json:"inbound,omitempty"`
}
