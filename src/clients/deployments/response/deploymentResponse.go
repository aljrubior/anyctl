package response

type DeploymentResponse struct {
	Id                    string      `json:"id"`
	Name                  string      `json:"name"`
	Target                Target      `json:"target"`
	Application           Application `json:"application"`
	Status                string      `json:"status"`
	DesiredVersion        string      `json:"desiredVersion"`
	LastSuccessfulVersion string      `json:"lastSuccessfulVersion"`
	Replicas              []Replica   `json:"replicas"`
}
