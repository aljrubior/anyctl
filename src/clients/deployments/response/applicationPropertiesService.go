package response

type ApplicationPropertiesService struct {
	ApplicationName  string            `json:"applicationName,omitempty"`
	Properties       map[string]string `json:"properties,omitempty"`
	SecureProperties map[string]string `json:"secureProperties,omitempty"`
}
