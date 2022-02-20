package managers

import (
	"github.com/aljrubior/anyctl/clients/deployments/response"
	"github.com/aljrubior/anyctl/errors"
	"github.com/aljrubior/anyctl/managers/entities"
	"github.com/aljrubior/anyctl/managers/requests"
	"github.com/aljrubior/anyctl/services"
	"strings"
)

func NewDefaultDeploymentManager(deploymentService services.DeploymentService, assetManager AssetManager) DefaultDeploymentManager {
	return DefaultDeploymentManager{
		deploymentService: deploymentService,
		assetManager:      assetManager,
	}
}

type DefaultDeploymentManager struct {
	deploymentService services.DeploymentService
	assetManager      AssetManager
}

func (this DefaultDeploymentManager) GetDeployments(ctx *entities.CurrentContextEntity) (*[]entities.DeploymentItemEntity, error) {

	resp, err := this.deploymentService.GetDeployments(ctx.OrganizationId, ctx.EnvironmentId, ctx.AuthorizationToken)

	if err != nil {
		return nil, err
	}

	return entities.NewDeploymentsEntityBuilder(resp).Build(), nil
}

func (this DefaultDeploymentManager) GetDeployment(ctx *entities.CurrentContextEntity, deploymentId string) (*entities.DeploymentEntity, error) {

	resp, err := this.deploymentService.GetDeployment(ctx.OrganizationId, ctx.EnvironmentId, ctx.AuthorizationToken, deploymentId)

	if err != nil {
		return nil, err
	}

	return entities.NewDeploymentEntityBuilder(resp).Build(), nil
}

func (this DefaultDeploymentManager) FindDeploymentByName(ctx *entities.CurrentContextEntity, deploymentName string) (*entities.DeploymentEntity, *[]entities.DeploymentItemEntity, error) {

	deployments, err := this.GetDeployments(ctx)

	if err != nil {
		return nil, nil, err
	}

	deploymentId := ""

	for _, v := range *deployments {
		if v.Name == deploymentName {
			deploymentId = v.Id
			break
		}
	}

	if deploymentId == "" {
		return nil, deployments, nil
	}

	deployment, err := this.GetDeployment(ctx, deploymentId)

	return deployment, deployments, err
}

func (this DefaultDeploymentManager) FindDeploymentContainsName(ctx *entities.CurrentContextEntity, deploymentName string) (*[]entities.DeploymentItemEntity, error) {

	deployments, err := this.GetDeployments(ctx)

	if err != nil {
		return nil, err
	}

	var deploymentsFound []entities.DeploymentItemEntity

	for _, v := range *deployments {
		if strings.Contains(v.Name, deploymentName) {
			deploymentsFound = append(deploymentsFound, v)
		}
	}

	return &deploymentsFound, err
}

func (this DefaultDeploymentManager) Deploy(ctx *entities.CurrentContextEntity, request *requests.DeploymentRequest, toEnvironment string) (*response.DeploymentResponse, error) {
	return this.deploymentService.Deploy(ctx.OrganizationId, toEnvironment, ctx.AuthorizationToken, request)
}

func (this DefaultDeploymentManager) StopDeployment(ctx *entities.CurrentContextEntity, deploymentId string) (*entities.DeploymentEntity, error) {
	request := requests.NewDeploymentStopRequestBuilder().Build()

	deployment, err := this.deploymentService.UpdateDeployment(ctx.OrganizationId, ctx.EnvironmentId, ctx.AuthorizationToken, deploymentId, request)

	if err != nil {
		return nil, err
	}

	return entities.NewDeploymentEntityBuilder(deployment).Build(), nil
}

func (this DefaultDeploymentManager) StartDeployment(ctx *entities.CurrentContextEntity, deploymentId string) (*entities.DeploymentEntity, error) {
	request := requests.NewDeploymentStartRequestBuilder().Build()

	deployment, err := this.deploymentService.UpdateDeployment(ctx.OrganizationId, ctx.EnvironmentId, ctx.AuthorizationToken, deploymentId, request)

	if err != nil {
		return nil, err
	}

	return entities.NewDeploymentEntityBuilder(deployment).Build(), nil
}

func (this DefaultDeploymentManager) DeleteDeployment(ctx *entities.CurrentContextEntity, deploymentId string) error {

	return this.deploymentService.DeleteDeployment(ctx.OrganizationId, ctx.EnvironmentId, ctx.AuthorizationToken, deploymentId)
}

func (this DefaultDeploymentManager) ScaleDeployment(ctx *entities.CurrentContextEntity, deploymentId string, desiredReplicas int) (*entities.DeploymentEntity, error) {

	d, err := this.GetDeployment(ctx, deploymentId)

	if err != nil {
		return nil, err
	}

	request := requests.NewDeploymentUpdateRequest(d.DeploymentResponse).WithReplicas(&desiredReplicas).Build()

	deployment, err := this.deploymentService.UpdateDeployment(ctx.OrganizationId, ctx.EnvironmentId, ctx.AuthorizationToken, deploymentId, request)

	if err != nil {
		return nil, err
	}

	return entities.NewDeploymentEntityBuilder(deployment).Build(), nil
}

func (this DefaultDeploymentManager) SetDeploymentAsset(ctx *entities.CurrentContextEntity, deploymentName, assetRef string) (*entities.DeploymentEntity, error) {

	deployment, _, err := this.FindDeploymentByName(ctx, deploymentName)

	if err != nil {
		return nil, err
	}

	asset, err := this.assetManager.FindAssetByRef(ctx, assetRef)

	if err != nil {
		return nil, this.ThrowAssetNotFoundError(assetRef, err)
	}

	if asset == nil {
		return nil, this.ThrowAssetNotFoundError(assetRef, nil)
	}

	request := requests.NewDeploymentUpdateRequest(deployment.DeploymentResponse).WithAsset(asset).Build()

	response, err := this.deploymentService.UpdateDeployment(ctx.OrganizationId, ctx.EnvironmentId, ctx.AuthorizationToken, deployment.Id, request)

	if err != nil {
		return nil, err
	}

	return entities.NewDeploymentEntityBuilder(response).Build(), nil
}

func (this DefaultDeploymentManager) ThrowAssetNotFoundError(assetName string, err error) error {
	return errors.NewAssetNotFoundError(assetName, nil).WithReason(err.Error())
}

func (this DefaultDeploymentManager) GetDeploymentSpecs(ctx *entities.CurrentContextEntity, deploymentId string) (*[]entities.DeploymentSpecEntity, error) {

	resp, err := this.deploymentService.GetDeploymentSpecs(ctx.OrganizationId, ctx.EnvironmentId, ctx.AuthorizationToken, deploymentId)

	if err != nil {
		return nil, err
	}

	return entities.NewDeploymentSpecEntitiesBuilder(resp).Build(), nil
}
