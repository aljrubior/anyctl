package handlers

type RuntimeFabricHandler interface {
	GetFabrics() error
	GetFabric(targetName string) error
	DescribeFabric(targetName string) error
	GetRuntimeFabricNodes(targetName string) error
	FindRuntimeFabricContainsName(targetName string) error
}
