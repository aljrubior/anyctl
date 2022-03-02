package response

import "github.com/aljrubior/anyctl/clients/privateSpaces/response"

type OrganizationPrivateSpaceResponse struct {
	Id                   string                   `json:"id"`
	Name                 string                   `json:"name"`
	Status               string                   `json:"status"`
	StatusMessage        string                   `json:"statusMessage"`
	Provisioning         PrivateSpaceProvisioning `json:"provisioning"`
	Region               string                   `json:"region"`
	OrganizationId       string                   `json:"organizationId"`
	ManagedFirewallRules []response.FirewallRules `json:"managedFirewallRules"`
	Environments         response.Environments    `json:"environments"`
	Network              PrivateSpaceNetwork      `json:"network"`
	FirewallRules        []response.FirewallRules `json:"firewallRules"`
}
