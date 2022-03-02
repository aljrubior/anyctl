package response

type Inbound struct {
	PublicUrl           string              `json:"publicUrl,omitempty"`
	PathRewrite         string              `json:"pathRewrite,omitempty"`
	DecoratedProperties DecoratedProperties `json:"decoratedProperties"`
}
