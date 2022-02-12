package response

type Inbound struct {
	PublicUrl           string               `yaml:"publicUrl,omitempty",json:"publicUrl,omitempty"`
	PathRewrite         string               `yaml:"pathRewrite,omitempty",json:"pathRewrite,omitempty"`
	DecoratedProperties *DecoratedProperties `yaml:"decoratedProperties",json:"decoratedProperties"`
}
