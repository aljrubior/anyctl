package handlers

type OrganizationPrivateSpaceHandler interface {
	GetPrivateSpaces() error
	GetPrivateSpace(psName string) error
	FindPrivateSpaceContainsName(targetName string) error
	DescribePrivateSpace(psName string) error
	GetFirewallRules(psName string) error
	GetNetwork(psName string) error
	GetFabrics(privateSpaceName string) error
}
