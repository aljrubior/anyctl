package response

type ApplicationPropertiesService struct {
	ApplicationName  string            `json:"applicationName"`
	Properties       map[string]string `json:"properties"`
	SecureProperties map[string]string `json:"secureProperties"`
}
