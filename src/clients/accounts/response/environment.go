package response

type Environment struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	OrganizationId string `json:"organizationId"`
	IsProduction   bool   `json:"isProduction"`
	Kind           string `json:"type"`
	ClientId       string `json:"clientId"`
}
