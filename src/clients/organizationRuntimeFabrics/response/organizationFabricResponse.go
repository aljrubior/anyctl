package response

import "github.com/aljrubior/anyctl/clients/fabrics/response"

type OrganizationFabricResponse struct {
	Id                        string                `yaml:"id",json:"id"`
	Name                      string                `yaml:"name",json:"name"`
	Region                    string                `yaml:"region",json:"region"`
	Vendor                    string                `yaml:"vendor",json:"vendor"`
	OrganizationId            string                `yaml:"organizationId",json:"organizationId"`
	Version                   string                `yaml:"version",json:"version"`
	Status                    string                `yaml:"status",json:"status"`
	DesiredVersion            string                `yaml:"desiredVersion",json:"desiredVersion"`
	AvailableUpgradeVersion   string                `yaml:"availableUpgradeVersion",json:"availableUpgradeVersion"`
	Nodes                     []response.FabricNode `yaml:"nodes",json:"nodes"`
	ClusterVersion            string                `yaml:"clusterVersion",json:"clusterVersion"`
	IsManaged                 bool                  `yaml:"isManaged",json:"isManaged"`
	Appliance                 response.Appliance    `yaml:"appliance",json:"appliance"`
	ClusterConfigurationLevel string                `yaml:"clusterConfigurationLevel",json:"clusterConfigurationLevel"`
	Features                  response.Features     `yaml:"features",json:"features"`
}
