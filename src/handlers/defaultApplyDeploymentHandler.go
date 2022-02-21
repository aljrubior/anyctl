package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/aljrubior/anyctl/errors"
	"github.com/aljrubior/anyctl/managers"
	"github.com/aljrubior/anyctl/managers/entities"
	"github.com/aljrubior/anyctl/managers/requests"
	"github.com/aljrubior/anyctl/manifests"
	"io/ioutil"
	"sigs.k8s.io/yaml"
)

func NewDefaultApplyDeploymentHandler(
	configManager managers.ConfigManager,
	deployerManager managers.DeployerManager,
	deploymentManager managers.DeploymentManager,
	targetManager managers.TargetManager,
	assetManager managers.AssetManager) DefaultApplyDeploymentHandler {

	return DefaultApplyDeploymentHandler{
		configManager:     configManager,
		deployerManager:   deployerManager,
		deploymentManager: deploymentManager,
		targetManager:     targetManager,
		assetManager:      assetManager,
	}
}

type DefaultApplyDeploymentHandler struct {
	configManager     managers.ConfigManager
	deployerManager   managers.DeployerManager
	deploymentManager managers.DeploymentManager
	targetManager     managers.TargetManager
	assetManager      managers.AssetManager
}

func (this DefaultApplyDeploymentHandler) Apply(filePath string) error {

	dataAsYaml, err := ioutil.ReadFile(filePath)

	if err != nil {
		return err
	}

	dataAsJson, err := yaml.YAMLToJSON(dataAsYaml)

	if err != nil {
		return err
	}

	var manifest manifests.DeploymentManifest

	err = json.Unmarshal(dataAsJson, &manifest)

	if err != nil {
		return err
	}

	ctx, err := this.configManager.GetCurrentContext()

	if err != nil {
		return err
	}

	targetId := manifest.Spec.Target.TargetId
	deploymentName := manifest.Spec.Name
	groupId := manifest.Spec.Application.Ref.GroupId
	artifactId := manifest.Spec.Application.Ref.ArtifactId
	version := manifest.Spec.Application.Ref.Version

	// Target Validation
	target, targets, err := this.targetManager.FindTargetById(ctx, targetId)

	if err != nil {
		return err
	}

	if target == nil {
		return this.ThrowTargetNotFoundError(targetId, targets)
	}

	// Ref Validation
	asset, err := this.assetManager.FindSpecificVersion(ctx, groupId, artifactId, version)

	if err != nil {
		return err
	}

	if asset == nil {
		return this.ThrowAssetNotFoundError(fmt.Sprintf("%s:%s:%s", groupId, artifactId, version))
	}

	deployment, _, err := this.deploymentManager.FindDeploymentByName(ctx, deploymentName)

	if deployment != nil {
		response := deployment.DeploymentResponse
		builder := requests.NewDeploymentApplyBuilder(response, manifest)
		request, err := builder.Apply()

		if err != nil {
			return err
		}

		this.deploymentManager.UpdateDeployment(ctx, deployment.Id, request)

		println(fmt.Sprintf("Deployment '%s' updated.", deploymentName))

		return nil
	}

	return nil
}

func (this DefaultApplyDeploymentHandler) ThrowTargetNotFoundError(targetId string, options *[]entities.TargetEntity) error {
	return errors.NewTargetNotFoundError(targetId, options)
}

func (this DefaultApplyDeploymentHandler) ThrowAssetNotFoundError(assetName string) error {
	return errors.NewAssetNotFoundError(assetName, nil)
}
