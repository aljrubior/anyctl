package response

type RuntimeVersion struct {
	BaseVersion string `json:"baseVersion"`
	Tag         string `json:"tag"`
	ReleaseDate int    `json:"releaseDate"`
	MinimumTag  string `json:"minimumTag"`
}
