package response

type PrivateSpaceResponse struct {
	Id                   string          `yaml:"id",json:"id"`
	Name                 string          `yaml:"name",json:"name"`
	Region               string          `yaml:"region",json:"region"`
	OrganizationId       string          `yaml:"organizationId",json:"organizationId"`
	Status               string          `yaml:"status",json:"status"`
	Version              string          `json:"version"`
	ManagedFirewallRules []FirewallRules `yaml:"managedFirewallRules",json:"managedFirewallRules"`
	Environments         Environments    `yaml:"environments"`
	Flavor               string          `json:"flavor"`
}
