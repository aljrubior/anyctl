package requests

type ArtifactRef struct {
	GroupId    string `json:"groupId,omitempty"`
	ArtifactId string `json:"artifactId,omitempty"`
	Version    string `json:"version,omitempty"`
	Packaging  string `json:"packaging,omitempty"`
}
