package response

type SharedSpaceResponse struct {
	Id                 string `json:"id",yaml:"id"`
	Name               string `json:"name",yaml:"name"`
	Region             string `json:"region"yaml:"region"`
	IsAdvertised       bool   `json:"isAdvertised"yaml:"isAdvertised"`
	RequiresPermission bool   `json:"requiresPermission"yaml:"requiresPermission"`
	Flavor             string `json:"flavor"yaml:"flavor"`
	PrivateSpaceId     string `json:"privateSpaceId"yaml:"privateSpaceId"`
	Status             string `json:"status"yaml:"status"`
	StatusMessage      string `json:"statusMessage"yaml:"statusMessage"`
}
