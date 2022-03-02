package response

type FabricTargetResponse struct {
	Id                        string    `yaml:"id",json:"id"`
	Name                      string    `yaml:"name",json:"name"`
	Version                   string    `yaml:"version",json:"version"`
	Type                      string    `yaml:"type",json:"type"`
	Runtimes                  []Runtime `yaml:"runtimes",json:"runtimes"`
	Status                    string    `yaml:"status",json:"status"`
	Environments              []string  `yaml:"environments",json:"environments"`
	IsAvailableForDeployments bool      `yaml:"isAvailableForDeployments",json:"isAvailableForDeployments"`
	Nodes                     []Node    `yaml:"nodes",json:"nodes"`
	Defaults                  Defaults  `yaml:"defaults",json:"defaults"`
	EnhancedSecurity          bool      `yaml:"enhancedSecurity",json:"enhancedSecurity"`
}
