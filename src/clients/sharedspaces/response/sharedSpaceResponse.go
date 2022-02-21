package response

type SharedSpaceResponse struct {
	Id                 string `json:"id"`
	Name               string `json:"name"`
	Region             string `json:"region"`
	IsAdvertised       bool   `json:"isAdvertised"`
	RequiresPermission bool   `json:"requiresPermission"`
	Flavor             string `json:"flavor"`
	PrivateSpaceId     string `json:"privateSpaceId"`
	Status             string `json:"status"`
	StatusMessage      string `json:"statusMessage"`
}
