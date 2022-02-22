package response

type FabricResponse struct {
	Id                        string       `json:"id"`
	Name                      string       `json:"name"`
	Region                    string       `json:"region"`
	Vendor                    string       `json:"vendor",`
	OrganizationId            string       `json:"organizationId"`
	Version                   string       `json:"version"`
	Status                    string       `json:"status"`
	ConsideredForScheduling   bool         `json:"consideredForScheduling"`
	DesiredVersion            string       `json:"desiredVersion"`
	AvailableUpgradeVersion   string       `json:"availableUpgradeVersion"`
	Nodes                     []FabricNode `json:"nodes"`
	SecondsSinceHeartbeat     int          `json:"secondsSinceHeartbeat"`
	StatusMessage             string       `json:"statusMessage"`
	KubernetesVersion         string       `json:"kubernetesVersion"`
	IsManaged                 bool         `json:"isManaged"`
	ClusterConfigurationLevel string       `json:"clusterConfigurationLevel"`
	DesiredInfraVersion       string       `json:"desiredInfraVersion"`
	InfraVersion              string       `json:"infraVersion"`
	InfraDeploymentId         string       `json:"infraDeploymentId"`
	LastInfraStatus           string       `json:"lastInfraStatus"`
	LastInfraMessage          string       `json:"lastInfraMessage"`
	Features                  Features     `json:"features"`
	NetworkId                 string       `json:"networkId"`
	OwnerId                   string       `json:"ownerId"`
	PartitionNumber           int          `json:"partitionNumber"`
	AcceptingTraffic          bool         `json:"acceptingTraffic"`
}
