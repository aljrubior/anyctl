package response

import "github.com/aljrubior/anyctl/clients/fabrics/response"

type OrganizationFabricResponse struct {
	Id                        string                `json:"id"`
	Name                      string                `json:"name"`
	Region                    string                `json:"region"`
	Vendor                    string                `json:"vendor"`
	OrganizationId            string                `json:"organizationId"`
	Version                   string                `json:"version"`
	Status                    string                `json:"status"`
	DesiredVersion            string                `json:"desiredVersion"`
	AvailableUpgradeVersion   string                `json:"availableUpgradeVersion"`
	Nodes                     []response.FabricNode `json:"nodes"`
	ClusterVersion            string                `json:"clusterVersion"`
	IsManaged                 bool                  `json:"isManaged"`
	Appliance                 response.Appliance    `json:"appliance"`
	ClusterConfigurationLevel string                `json:"clusterConfigurationLevel"`
	Features                  response.Features     `json:"features"`
}
