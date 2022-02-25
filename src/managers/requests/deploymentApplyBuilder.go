package requests

import (
	"encoding/json"
	"github.com/aljrubior/anyctl/clients/deployments/response"
	"github.com/aljrubior/anyctl/manifests"
	"reflect"
)

func NewDeploymentApplyBuilder(response *response.DeploymentResponse, manifest manifests.DeploymentManifest) DeploymentApplyBuilder {
	return DeploymentApplyBuilder{
		response: response,
		manifest: manifest,
	}
}

type DeploymentApplyBuilder struct {
	response *response.DeploymentResponse
	manifest manifests.DeploymentManifest
}

func (this *DeploymentApplyBuilder) Apply() (DeploymentRequest, error) {

	if this.response == nil {
		return this.buildRequest(this.manifest)
	}

	this.response.Target = this.applyTarget(this.response.Target, this.manifest.Spec.Target)
	this.response.Application = this.applyApplication(this.response.Application, this.manifest.Spec.Application)

	data, err := json.Marshal(this.response)

	if err != nil {
		return DeploymentRequest{}, err
	}

	var request DeploymentRequest

	err = json.Unmarshal(data, &request)

	if err != nil {
		return DeploymentRequest{}, err
	}

	return request, nil
}

func (this *DeploymentApplyBuilder) buildRequest(manifest manifests.DeploymentManifest) (DeploymentRequest, error) {

	data, err := json.Marshal(manifest.Spec)

	if err != nil {
		return DeploymentRequest{}, err
	}

	var request DeploymentRequest

	err = json.Unmarshal(data, &request)

	if err != nil {
		return DeploymentRequest{}, err
	}

	data, _ = json.Marshal(request)

	if request.Application.Configuration.ApplicationPropertiesService.Properties == nil {
		request.Application.Configuration.ApplicationPropertiesService.Properties = map[string]string{}
	}

	if request.Application.Configuration.ApplicationPropertiesService.SecureProperties == nil {
		request.Application.Configuration.ApplicationPropertiesService.SecureProperties = map[string]string{}
	}
	return request, err
}

func (this *DeploymentApplyBuilder) applyTarget(target, withTarget response.Target) response.Target {

	target.DeploymentSettings = this.applyDeploymentSettings(target.DeploymentSettings, withTarget.DeploymentSettings)

	if target.Replicas != withTarget.Replicas {
		target.Replicas = withTarget.Replicas
	}

	return target
}

func (this *DeploymentApplyBuilder) applyDeploymentSettings(settings, withSettings response.DeploymentSettings) response.DeploymentSettings {

	settings.Resources = this.applyResources(settings.Resources, withSettings.Resources)

	if settings.Clustered != withSettings.Clustered {
		settings.Clustered = withSettings.Clustered
	}

	if settings.EnforceDeployingReplicasAcrossNodes != withSettings.EnforceDeployingReplicasAcrossNodes {
		settings.EnforceDeployingReplicasAcrossNodes = withSettings.EnforceDeployingReplicasAcrossNodes
	}

	settings.Http = this.applyHttp(settings.Http, withSettings.Http)
	settings.Jvm = this.applyJvm(settings.Jvm, withSettings.Jvm)

	if settings.RuntimeVersion != withSettings.RuntimeVersion {
		settings.RuntimeVersion = withSettings.RuntimeVersion
	}

	if settings.LastMileSecurity != withSettings.LastMileSecurity {
		settings.LastMileSecurity = withSettings.LastMileSecurity
	}

	if settings.ForwardSslSession != withSettings.ForwardSslSession {
		settings.ForwardSslSession = withSettings.ForwardSslSession
	}

	if settings.UpdateStrategy != withSettings.UpdateStrategy {
		settings.UpdateStrategy = withSettings.UpdateStrategy
	}

	return settings
}

func (this *DeploymentApplyBuilder) applyJvm(jvm, withJvm response.Jvm) response.Jvm {

	if withJvm.Args != jvm.Args {
		jvm.Args = withJvm.Args
	}

	return jvm
}

func (this *DeploymentApplyBuilder) applyResources(resources, withResources response.Resources) response.Resources {

	if withResources.Cpu.Reserved != resources.Cpu.Reserved {
		resources.Cpu.Reserved = withResources.Cpu.Reserved
	}

	if withResources.Cpu.Limit != resources.Cpu.Limit {
		resources.Cpu.Limit = withResources.Cpu.Limit
	}

	if withResources.Memory.Reserved != resources.Memory.Reserved {
		resources.Memory.Reserved = withResources.Memory.Reserved
	}

	if withResources.Memory.Limit != resources.Memory.Limit {
		resources.Memory.Limit = withResources.Memory.Limit
	}

	return resources
}

func (this *DeploymentApplyBuilder) applyHttp(http, withHttp response.Http) response.Http {

	if http.Inbound.PublicUrl != withHttp.Inbound.PublicUrl {
		http.Inbound.PublicUrl = withHttp.Inbound.PublicUrl
	}

	if http.Inbound.PathRewrite != withHttp.Inbound.PathRewrite {
		http.Inbound.PathRewrite = withHttp.Inbound.PathRewrite
	}

	return http
}

func (this *DeploymentApplyBuilder) applyApplication(application, withApplication response.Application) response.Application {

	application.Ref = this.applyRef(application.Ref, withApplication.Ref)

	application.Configuration = this.applyConfiguration(application.Configuration, withApplication.Configuration)

	if application.DesiredState != withApplication.DesiredState {
		application.DesiredState = withApplication.DesiredState
	}

	return application
}

func (this *DeploymentApplyBuilder) applyRef(ref, withRef response.Asset) response.Asset {

	if ref.GroupId != withRef.GroupId {
		ref.GroupId = withRef.GroupId
	}

	if ref.ArtifactId != withRef.ArtifactId {
		ref.ArtifactId = withRef.ArtifactId
	}

	if ref.Version != withRef.Version {
		ref.Version = withRef.Version
	}

	if ref.Packaging != withRef.Packaging {
		ref.Packaging = withRef.Packaging
	}

	return ref
}

func (this *DeploymentApplyBuilder) applyConfiguration(configuration, withConfiguration response.ApplicationConfiguration) response.ApplicationConfiguration {

	if !reflect.DeepEqual(configuration.ApplicationPropertiesService.Properties, withConfiguration.ApplicationPropertiesService.Properties) {
		configuration.ApplicationPropertiesService.Properties = withConfiguration.ApplicationPropertiesService.Properties
	}

	if !reflect.DeepEqual(configuration.ApplicationPropertiesService.SecureProperties, withConfiguration.ApplicationPropertiesService.SecureProperties) {
		configuration.ApplicationPropertiesService.SecureProperties = withConfiguration.ApplicationPropertiesService.SecureProperties
	}

	if configuration.ApplicationPropertiesService.Properties == nil {
		configuration.ApplicationPropertiesService.Properties = map[string]string{}
	}

	if configuration.ApplicationPropertiesService.SecureProperties == nil {
		configuration.ApplicationPropertiesService.SecureProperties = map[string]string{}
	}

	return configuration
}
