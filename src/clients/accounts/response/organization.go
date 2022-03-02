package response

type Organization struct {
	Name               string        `json:"name"`
	Id                 string        `json:"id"`
	SubOrganizationIds []string      `json:"subOrganizationIds"`
	Environments       []Environment `json:"environments"`
}
