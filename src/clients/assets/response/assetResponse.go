package response

type AssetResponse struct {
	GroupId           string    `json:"groupId"`
	AssetId           string    `json:"assetId"`
	Version           string    `json:"version"`
	Description       string    `json:"description"`
	VersionGroup      string    `json:"versionGroup"`
	ProductAPIVersion string    `json:"productAPIVersion"`
	IsPublic          bool      `json:"isPublic"`
	Name              string    `json:"name"`
	Type              string    `json:"type"`
	IsSnapshot        bool      `json:"isSnapshot"`
	Status            string    `json:"status"`
	AssetLink         string    `json:"assetLink"`
	CreatedAt         string    `json:"createdAt"`
	UpdatedAt         string    `json:"updatedAt"`
	RuntimeVersion    string    `json:"runtimeVersion"`
	CreatedBy         CreatedBy `json:"createdBy"`
}
