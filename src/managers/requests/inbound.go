package requests

type Inbound struct {
	PublicUrl   string `json:"publicUrl"`
	PathRewrite string `json:"pathRewrite,omitempty"`
}
