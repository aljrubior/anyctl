package response

type FabricResponse struct {
	Id                        string       `yaml:"id",json:"id"`
	Name                      string       `yaml:"name"json:"name"`
	Region                    string       `yaml:"region",json:"region"`
	Vendor                    string       `yaml:"vendor",json:"vendor",`
	OrganizationId            string       `yaml:"organizationId",json:"organizationId"`
	Version                   string       `yaml:"version",json:"version"`
	Status                    string       `yaml:"status",json:"status"`
	ConsideredForScheduling   bool         `yaml:"consideredForScheduling",json:"consideredForScheduling"`
	DesiredVersion            string       `yaml:"desiredVersion",json:"desiredVersion"`
	AvailableUpgradeVersion   string       `yaml:"availableUpgradeVersion",json:"availableUpgradeVersion"`
	Nodes                     []FabricNode `yaml:"nodes",json:"nodes"`
	SecondsSinceHeartbeat     int          `yaml:"secondsSinceHeartbeat",json:"secondsSinceHeartbeat"`
	StatusMessage             string       `yaml:"statusMessage",json:"statusMessage"`
	KubernetesVersion         string       `yaml:"kubernetesVersion",json:"kubernetesVersion"`
	IsManaged                 bool         `yaml:"isManaged",json:"isManaged"`
	ClusterConfigurationLevel string       `yaml:"clusterConfigurationLevel",json:"clusterConfigurationLevel"`
	DesiredInfraVersion       string       `yaml:"desiredInfraVersion",json:"desiredInfraVersion"`
	InfraVersion              string       `yaml:"infraVersion",json:"infraVersion"`
	InfraDeploymentId         string       `yaml:"infraDeploymentId",json:"infraDeploymentId"`
	LastInfraStatus           string       `yaml:"lastInfraStatus",json:"lastInfraStatus"`
	LastInfraMessage          string       `yaml:"lastInfraMessage",json:"lastInfraMessage"`
	Features                  Features     `yaml:"features",json:"features"`
	NetworkId                 string       `yaml:"networkId",json:"networkId"`
	OwnerId                   string       `yaml:"ownerId",json:"ownerId"`
	PartitionNumber           int          `yaml:"partitionNumber",json:"partitionNumber"`
	AcceptingTraffic          bool         `yaml:"acceptingTraffic",json:"acceptingTraffic"`
}
