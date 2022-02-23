package response

type Asset struct {
	GroupId    string `json:"groupId"`
	ArtifactId string `json:"artifactId"`
	Version    string `json:"version"`
	Packaging  string `json:"packaging"`
}
