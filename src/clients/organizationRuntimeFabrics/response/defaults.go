package response

type Defaults struct {
	ApiQueryApplication ApiQueryApplicationDefault `yaml:"apiQueryApplication",json:"apiQueryApplication"`
	MuleApplication     MuleApplicationDefault     `yaml:"muleApplication",json:"muleApplication"`
	EnhancedSecurity    bool                       `yaml:"enhancedSecurity",json:"enhancedSecurity"`
}
