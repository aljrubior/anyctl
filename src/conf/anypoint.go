package conf

type Anypoint struct {
	Protocol  string            `yaml:"protocol"`
	Host      string            `yaml:"host"`
	Port      int               `yaml:"port"`
	Resources AnypointResources `yaml:"resources"`
}
