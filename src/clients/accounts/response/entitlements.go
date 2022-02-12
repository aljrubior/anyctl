package response

type Entitlements struct {
	VCoresProduction ResourceUsage `json:"vCoresProduction"`
	VCoresSandbox    ResourceUsage `json:"vCoresSandbox"`
	VCoresDesign     ResourceUsage `json:"vCoresDesign"`
	StaticIps        ResourceUsage `json:"staticIps"`
	Vpcs             ResourceUsage `json:"vpcs"`
	Vpns             ResourceUsage `json:"vpns"`
	LoadBalancer     ResourceUsage `json:"loadBalancer"`
}
