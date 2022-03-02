package managers

import (
	"encoding/json"
	"fmt"
	"github.com/aljrubior/anyctl/clients/deployments/response"
	"github.com/aljrubior/anyctl/errors"
	"github.com/aljrubior/anyctl/managers/entities"
	"github.com/aljrubior/anyctl/managers/requests"
	"github.com/aljrubior/anyctl/managers/wrappers"
	"strings"
)

func NewDefaultDeployerManager(
	deploymentManager DeploymentManager,
	assetManager AssetManager,
	organizationRuntimeFabricManager OrganizationRuntimeFabricManager) DefaultDeployerManager {
	return DefaultDeployerManager{
		deploymentManager:                deploymentManager,
		assetManager:                     assetManager,
		organizationRuntimeFabricManager: organizationRuntimeFabricManager,
	}
}

type DefaultDeployerManager struct {
	DeployerManager
	deploymentManager                DeploymentManager
	assetManager                     AssetManager
	organizationRuntimeFabricManager OrganizationRuntimeFabricManager
}

func (this DefaultDeployerManager) Deploy(ctx *entities.CurrentContextEntity, deploymentName, assetRef, targetName, runtimeBaseVersion string) (*entities.DeploymentEntity, error) {

	asset, err := this.assetManager.FindAssetByRef(ctx, assetRef)

	if err != nil {
		return nil, this.ThrowAssetNotFoundError(assetRef, err)
	}

	if asset == nil {
		return nil, this.ThrowAssetNotFoundError(assetRef, nil)
	}

	target, targets, err := this.organizationRuntimeFabricManager.FindFabricTargetByName(ctx, targetName)

	if err != nil {
		return nil, err
	}

	if target == nil {
		return nil, this.ThrowFabricTargetNotFoundError(targetName, targets)
	}

	versionRef := wrappers.NewFabricTargetEntityWrapper(target).GetRuntimeVersionRef("mule", runtimeBaseVersion)

	spec := requests.
		NewApplicationRequestSpecBuilder().
		WithApplicationName(deploymentName).
		WithTarget("MC", target.Id).
		WithRuntimeVersion(versionRef).
		WithAsset(asset.GroupId, asset.Name, asset.Version, "jar").
		Build()

	request := requests.NewApplicationRequestBuilder(spec).Build()

	resp, err := this.deploymentManager.Deploy(ctx, request, ctx.EnvironmentId)

	if err != nil {
		return nil, err
	}

	return entities.NewDeploymentEntityBuilder(resp).Build(), nil
}

func (this DefaultDeployerManager) CopyDeployment(ctx *entities.CurrentContextEntity, deployment *entities.DeploymentEntity, withName string, toTarget *entities.TargetEntity, toEnvironmentId string) (*entities.DeploymentEntity, error) {

	request, err := this.clone(deployment.DeploymentResponse, withName)

	if err != nil {
		return nil, err
	}

	if toTarget != nil {
		request.Target.TargetId = (*toTarget).GetId()
	}

	resp, err := this.deploymentManager.Deploy(ctx, request, toEnvironmentId)

	if err != nil {
		return nil, err
	}

	return entities.NewDeploymentEntityBuilder(resp).Build(), nil
}

func (this DefaultDeployerManager) clone(response response.DeploymentResponse, withName string) (*requests.DeploymentRequest, error) {

	currentName := response.Name

	var request requests.DeploymentRequest

	data, err := json.Marshal(response)

	if err != nil {
		return nil, err
	}

	if withName == "" {
		if err := json.Unmarshal(data, &request); err != nil {
			return nil, err
		}
	} else {
		responseAsString := string(data)
		responseAfterReplace := strings.ReplaceAll(responseAsString, currentName, withName)

		if err := json.Unmarshal([]byte(responseAfterReplace), &request); err != nil {
			return nil, err
		}
	}

	if fmt.Sprintf("%v", request.Application.Configuration) == "{}" {
		request.Application.Configuration = this.buildDefaultConfigurations(request.Name)
	}
	return &request, nil

}

func (this DefaultDeployerManager) buildDefaultConfigurations(applicationName string) *requests.ApplicationConfiguration {
	return &requests.ApplicationConfiguration{
		ApplicationPropertiesService: requests.NewApplicationPropertiesService(applicationName, nil, nil),
	}
}

func (this DefaultDeployerManager) ThrowAssetNotFoundError(assetName string, err error) error {
	return errors.NewAssetNotFoundError(assetName, nil).WithReason(err.Error())
}

func (this DefaultDeployerManager) ThrowFabricTargetNotFoundError(targetName string, options *[]entities.FabricTargetEntity) error {
	return errors.NewFabricTargetNotFoundError(targetName, options)
}
