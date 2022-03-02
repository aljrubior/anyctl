package handlers

import (
	"fmt"
	"github.com/aljrubior/anyctl/managers"
)

func NewDefaultRunHandler(configManager managers.ConfigManager, deployerManager managers.DeployerManager) *DefaultRunHandler {
	return &DefaultRunHandler{
		configManager:   configManager,
		deployerManager: deployerManager,
	}

}

type DefaultRunHandler struct {
	RunHandler
	configManager   managers.ConfigManager
	deployerManager managers.DeployerManager
}

func (this DefaultRunHandler) Deploy(deploymentName, assetRef, targetName, runtimeVersion string) error {

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	_, err = this.deployerManager.Deploy(ctx, deploymentName, assetRef, targetName, runtimeVersion)

	if err != nil {
		return err
	}

	println(fmt.Sprintf("Deployment '%s' created.", deploymentName))

	return nil
}
