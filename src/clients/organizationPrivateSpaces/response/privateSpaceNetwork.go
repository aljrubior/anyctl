package response

type PrivateSpaceNetwork struct {
	Region            string   `yaml:"region",json:"region"`
	CidrBlock         string   `yaml:"cidrBlock",json:"cidrBlock"`
	InboundStaticIps  []string `yaml:"inboundStaticIps",json:"inboundStaticIps"`
	OutboundStaticIps []string `yaml:"outboundStaticIps",json:"outboundStaticIps"`
	DnsTarget         string   `yaml:"dnsTarget",json:"dnsTarget"`
}
