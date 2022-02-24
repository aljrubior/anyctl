package requests

func NewApplicationRequestBuilder(fromSpec *ApplicationRequestSpec) *ApplicationRequestBuilder {
	return &ApplicationRequestBuilder{
		spec: fromSpec,
	}
}

type ApplicationRequestBuilder struct {
	spec *ApplicationRequestSpec
}

func (this ApplicationRequestBuilder) Build() *DeploymentRequest {

	return &DeploymentRequest{
		Name:   this.spec.ApplicationName,
		Labels: []string{"beta"},
		Target: &Target{
			Provider: this.spec.TargetProvider,
			TargetId: this.spec.TargetId,
			DeploymentSettings: &DeploymentSettings{
				Resources: Resources{
					Cpu: ResourceItem{
						Reserved: &this.spec.CpuReserved,
						Limit:    &this.spec.CpuLimit,
					},
					Memory: ResourceItem{
						Reserved: &this.spec.MemoryReserved,
						Limit:    &this.spec.MemoryLimit,
					},
				},
				Clustered:                           this.spec.Clustered,
				EnforceDeployingReplicasAcrossNodes: this.spec.EnforceDeployingReplicasAcrossNodes,
				Http: Http{
					Inbound: Inbound{
						PublicUrl:   this.spec.PublicUrl,
						PathRewrite: this.spec.PathRewrite,
					},
				},
				Jvm:               Jvm{},
				RuntimeVersion:    this.spec.RuntimeVersion,
				LastMileSecurity:  this.spec.LastMileSecurity,
				ForwardSslSession: this.spec.ForwardSslSession,
				UpdateStrategy:    this.spec.UpdateStrategy,
			},
			Replicas: this.spec.Replicas,
		},
		Application: Application{
			Ref: &ArtifactRef{
				GroupId:    this.spec.GroupId,
				ArtifactId: this.spec.ArtifactId,
				Version:    this.spec.ArtifactVersion,
				Packaging:  this.spec.Packaging,
			},
			Assets:       this.spec.Assets,
			DesiredState: this.spec.DesiredState,
			Configuration: &ApplicationConfiguration{
				ApplicationPropertiesService: ApplicationPropertiesService{
					ApplicationName:  this.spec.ApplicationName,
					Properties:       map[string]string{},
					SecureProperties: map[string]string{},
				},
			},
		},
	}
}
