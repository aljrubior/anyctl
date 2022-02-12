package response

type Asset struct {
	GroupId    string `yaml:"groupId",json:"groupId"`
	ArtifactId string `yaml:"artifactId",json:"artifactId"`
	Version    string `yaml:"version",json:"version"`
	Packaging  string `yaml:"packaging",json:"packaging"`
}
