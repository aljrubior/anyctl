package response

type FirewallRules struct {
	CidrBlock string `json:"cidrBlock"`
	Protocol  string `json:"protocol"`
	FromPort  int    `json:"fromPort"`
	ToPort    int    `json:"toPort"`
	Type      string `json:"type"`
}
