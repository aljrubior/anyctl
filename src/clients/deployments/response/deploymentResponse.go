package response

type DeploymentResponse struct {
	Id                    string      `yaml:"id",json:"id"`
	Name                  string      `yaml:"name",json:"name"`
	Target                Target      `yaml:"target",json:"target"`
	Application           Application `yaml:"application",json:"application"`
	Status                string      `yaml:"status",json:"status"`
	DesiredVersion        string      `yaml:"desiredVersion",json:"desiredVersion"`
	LastSuccessfulVersion string      `yaml:"lastSuccessfulVersion",json:"lastSuccessfulVersion"`
	Replicas              []Replica   `yaml:"replicas",json:"replicas"`
}
