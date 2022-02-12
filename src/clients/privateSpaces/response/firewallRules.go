package response

type FirewallRules struct {
	CidrBlock string `yaml:"cidrBlock",json:"cidrBlock"`
	Protocol  string `yaml:"protocol",json:"protocol"`
	FromPort  int    `yaml:"fromPort",json:"fromPort"`
	ToPort    int    `yaml:"toPort",json:"toPort"`
	Type      string `yaml:"type",json:"type"`
}
