package response

type PrivateSpaceNetwork struct {
	Region            string   `json:"region"`
	CidrBlock         string   `json:"cidrBlock"`
	InboundStaticIps  []string `json:"inboundStaticIps"`
	OutboundStaticIps []string `json:"outboundStaticIps"`
	DnsTarget         string   `json:"dnsTarget"`
}
