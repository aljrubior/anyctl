package response

import "github.com/aljrubior/anyctl/clients/privateSpaces/response"

type OrganizationPrivateSpaceResponse struct {
	Id                   string                   `yaml:"id",json:"id"`
	Name                 string                   `yaml:"name",json:"name"`
	Status               string                   `yaml:"status",json:"status"`
	StatusMessage        string                   `yaml:"statusMessage",json:"statusMessage"`
	Provisioning         PrivateSpaceProvisioning `yaml:"provisioning",json:"provisioning"`
	Region               string                   `yaml:"region",json:"region"`
	OrganizationId       string                   `yaml:"organizationId",json:"organizationId"`
	ManagedFirewallRules []response.FirewallRules `yaml:"managedFirewallRules",json:"managedFirewallRules"`
	Environments         response.Environments    `yaml:"environments"`
	Network              PrivateSpaceNetwork      `yaml:"network",json:"network"`
	FirewallRules        []response.FirewallRules `yaml:"firewallRules",json:"firewallRules"`
}
