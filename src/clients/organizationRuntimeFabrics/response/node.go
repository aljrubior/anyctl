package response

type Node struct {
	Id                        string `yaml:"id",json:"id"`
	Location                  string `yaml:"location",json:"location"`
	IsAvailableForDeployments bool   `yaml:"isAvailableForDeployments",json:"isAvailableForDeployments"`
}
