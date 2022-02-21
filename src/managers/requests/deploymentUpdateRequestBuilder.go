package requests

import (
	"github.com/aljrubior/anyctl/clients/deployments/response"
	"github.com/aljrubior/anyctl/managers/entities"
)

func NewDeploymentUpdateRequest(response response.DeploymentResponse) *DeploymentUpdateRequestBuilder {

	return &DeploymentUpdateRequestBuilder{
		response: response,
	}
}

type DeploymentUpdateRequestBuilder struct {
	response        response.DeploymentResponse
	desiredReplicas *int
	assetEntity     *entities.AssetEntity
}

func (this *DeploymentUpdateRequestBuilder) WithReplicas(desiredReplicas *int) *DeploymentUpdateRequestBuilder {
	this.desiredReplicas = desiredReplicas
	return this
}

func (this *DeploymentUpdateRequestBuilder) WithAsset(assetEntity *entities.AssetEntity) *DeploymentUpdateRequestBuilder {
	this.assetEntity = assetEntity
	return this
}

func (this DeploymentUpdateRequestBuilder) Build() *DeploymentRequest {

	return &DeploymentRequest{
		Name: this.response.Name,
		Target: Target{
			Provider: this.response.Target.Provider,
			TargetId: this.response.Target.TargetId,
			Replicas: this.getReplicas(),
			Type:     this.response.Target.Type,
		},
		Application: Application{
			Ref:          this.buildArtifactRef(),
			DesiredState: this.response.Application.DesiredState,
		},
	}
}

func (this DeploymentUpdateRequestBuilder) buildArtifactRef() ArtifactRef {
	if this.assetEntity == nil {
		return ArtifactRef{
			GroupId:    this.response.Application.Ref.GroupId,
			ArtifactId: this.response.Application.Ref.ArtifactId,
			Version:    this.response.Application.Ref.Version,
			Packaging:  this.response.Application.Ref.Packaging,
		}
	}

	return ArtifactRef{
		GroupId:    this.assetEntity.GroupId,
		ArtifactId: this.assetEntity.AssetId,
		Version:    this.assetEntity.Version,
		Packaging:  "jar",
	}
}

func (this DeploymentUpdateRequestBuilder) getReplicas() int {
	if this.desiredReplicas == nil {
		return this.response.Target.Replicas
	}

	return *this.desiredReplicas
}
