package requests

type Inbound struct {
	PublicUrl   *string `json:"publicUrl,omitempty"`
	PathRewrite *string `json:"pathRewrite,omitempty"`
}
