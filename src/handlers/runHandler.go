package handlers

type RunHandler interface {
	Deploy(name, assetRef, targetName, runtimeVersion string) error
}
