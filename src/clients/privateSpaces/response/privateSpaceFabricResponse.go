package response

import "github.com/aljrubior/anyctl/clients/fabrics/response"

type PrivateSpaceFabricResponse struct {
	Id                        string                `json:"id"`
	Name                      string                `json:"name"`
	Region                    string                `json:"region"`
	Vendor                    string                `json:"vendor"`
	OrganizationId            string                `json:"organizationId"`
	Version                   string                `json:"version"`
	Status                    string                `json:"status"`
	ConsideredForScheduling   bool                  `json:"consideredForScheduling"`
	DesiredVersion            string                `json:"desiredVersion"`
	AvailableUpgradeVersion   string                `json:"availableUpgradeVersion"`
	Nodes                     []response.FabricNode `json:"nodes"`
	LastUpgradeTimestamp      int64                 `json:"lastUpgradeTimestamp"`
	SecondsSinceHeartbeat     int                   `json:"secondsSinceHeartbeat"`
	KubernetesVersion         string                `json:"kubernetesVersion"`
	IsManaged                 bool                  `json:"isManaged"`
	Appliance                 response.Appliance    `json:"appliance"`
	ClusterConfigurationLevel string                `json:"clusterConfigurationLevel"`
	DesiredInfraVersion       string                `json:"desiredInfraVersion"`
	Features                  response.Features     `json:"features"`
	InfraVersion              string                `json:"infraVersion"`
	InfraDeploymentId         string                `json:"infraDeploymentId"`
	LastInfraStatus           string                `json:"lastInfraStatus"`
	NetworkId                 string                `json:"networkId"`
	OwnerId                   string                `json:"ownerId"`
	CreatedAt                 int64                 `json:"createdAt"`
	PartitionNumber           int                   `json:"partitionNumber"`
	AcceptingTraffic          bool                  `json:"acceptingTraffic"`
}
