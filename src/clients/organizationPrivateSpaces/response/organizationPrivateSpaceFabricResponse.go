package response

type OrganizationPrivateSpaceFabricResponse struct {
	Id                    string `json:"id"`
	Name                  string `json:"name"`
	Region                string `json:"region"`
	Status                string `json:"status"`
	LastUpgradeTimestamp  int64  `json:"lastUpgradeTimestamp"`
	SecondsSinceHeartbeat int    `json:"secondsSinceHeartbeat"`
	StatusMessage         string `json:"statusMessage"`
}
