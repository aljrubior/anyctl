package requests

type ApplicationPropertiesService struct {
	ApplicationName  string                `json:"applicationName,omitempty"`
	Properties       ApplicationProperties `json:"properties,omitempty"`
	SecureProperties SecureProperties      `json:"secureProperties,omitempty"`
}
