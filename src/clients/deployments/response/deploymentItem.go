package response

type DeploymentItem struct {
	Id          string      `json:"id"`
	Name        string      `json:"name"`
	Target      Target      `json:"target"`
	Status      string      `json:"status"`
	Application Application `json:"application"`
}
