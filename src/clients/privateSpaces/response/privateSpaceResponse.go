package response

type PrivateSpaceResponse struct {
	Id                   string          `json:"id"`
	Name                 string          `json:"name"`
	Region               string          `json:"region"`
	OrganizationId       string          `json:"organizationId"`
	Status               string          `json:"status"`
	Version              string          `json:"version"`
	ManagedFirewallRules []FirewallRules `json:"managedFirewallRules"`
	Environments         Environments    `json:"environments"`
	Flavor               string          `json:"flavor"`
}
