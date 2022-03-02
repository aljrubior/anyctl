package requests

func NewApplicationPropertiesService(applicationName string, properties, secureProperties *map[string]string) ApplicationPropertiesService {

	var props, secureProps map[string]string

	if properties == nil {
		props = map[string]string{}
	}

	if secureProperties == nil {
		secureProps = map[string]string{}
	}

	return ApplicationPropertiesService{
		ApplicationName:  applicationName,
		Properties:       props,
		SecureProperties: secureProps,
	}
}

type ApplicationPropertiesService struct {
	ApplicationName  string            `json:"applicationName"`
	Properties       map[string]string `json:"properties"`
	SecureProperties map[string]string `json:"secureProperties"`
}
