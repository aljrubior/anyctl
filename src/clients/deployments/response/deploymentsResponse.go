package response

type DeploymentsResponse struct {
	Total int              `json:"total"`
	Items []DeploymentItem `json:"items"`
}
