package response

type OrganizationResponse struct {
	Name                  string        `json:"name"`
	Id                    string        `json:"id"`
	OwnerId               string        `json:"ownerId"`
	ClientId              string        `json:"clientId"`
	ParentOrganizationIds []string      `json:"parentOrganizationIds"`
	SubOrganizationIds    []string      `json:"subOrganizationIds"`
	Entitlements          Entitlements  `json:"entitlements"`
	IsMaster              bool          `json:"isMaster"`
	Environments          []Environment `json:"environments"`
	Subscription          Subscription  `json:"subscription"`
}
