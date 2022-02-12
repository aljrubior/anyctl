package requests

func NewApplicationRequestSpecBuilder() *ApplicationRequestSpecBuilder {
	return &ApplicationRequestSpecBuilder{
		applicationRequestSpec: ApplicationRequestSpec{
			ApplicationName:                     "",
			Labels:                              []string{"beta"},
			TargetProvider:                      "",
			TargetId:                            "",
			CpuReserved:                         "20m",
			CpuLimit:                            "16000m",
			MemoryReserved:                      "700Mi",
			MemoryLimit:                         "700Mi",
			Clustered:                           false,
			EnforceDeployingReplicasAcrossNodes: false,
			PublicUrl:                           nil,
			PathRewrite:                         nil,
			RuntimeVersion:                      "",
			LastMileSecurity:                    false,
			ForwardSslSession:                   false,
			UpdateStrategy:                      "rolling",
			Replicas:                            1,
			GroupId:                             "",
			ArtifactId:                          "",
			ArtifactVersion:                     "",
			Packaging:                           "",
			Assets:                              make([]interface{}, 0),
			DesiredState:                        "STARTED",
		},
	}
}

type ApplicationRequestSpecBuilder struct {
	applicationRequestSpec ApplicationRequestSpec
}

func (this *ApplicationRequestSpecBuilder) WithApplicationName(name string) *ApplicationRequestSpecBuilder {
	this.applicationRequestSpec.ApplicationName = name
	return this
}

func (this *ApplicationRequestSpecBuilder) WithTarget(provider, id string) *ApplicationRequestSpecBuilder {
	this.applicationRequestSpec.TargetProvider = provider
	this.applicationRequestSpec.TargetId = id
	return this
}

func (this *ApplicationRequestSpecBuilder) WithRuntimeVersion(version string) *ApplicationRequestSpecBuilder {
	this.applicationRequestSpec.RuntimeVersion = version
	return this
}

func (this *ApplicationRequestSpecBuilder) WithReplicas(total int) *ApplicationRequestSpecBuilder {
	this.applicationRequestSpec.Replicas = total
	return this
}

func (this *ApplicationRequestSpecBuilder) WithAsset(groupId, artifactId, version, packaging string) *ApplicationRequestSpecBuilder {
	this.applicationRequestSpec.GroupId = groupId
	this.applicationRequestSpec.ArtifactId = artifactId
	this.applicationRequestSpec.ArtifactVersion = version
	this.applicationRequestSpec.Packaging = packaging
	return this
}

func (this *ApplicationRequestSpecBuilder) Build() *ApplicationRequestSpec {
	return &this.applicationRequestSpec

}
