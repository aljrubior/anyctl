package handlers

type PrivateSpaceHandler interface {
	GetPrivateSpaces(privateSpaceId string) error
	GetPrivateSpace(privateSpaceName string) error
	DescribePrivateSpace(privateSpaceName string) error

	GetManagedFirewallRules(privateSpaceName string) error

	GetFabrics(privateSpaceId string) error
}
