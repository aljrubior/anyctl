package requests

type ApplicationRequestSpec struct {
	ApplicationName                     string
	Labels                              []string
	TargetProvider                      string
	TargetId                            string
	CpuReserved                         string
	CpuLimit                            string
	MemoryReserved                      string
	MemoryLimit                         string
	Clustered                           bool
	EnforceDeployingReplicasAcrossNodes bool
	PublicUrl                           string
	PathRewrite                         string
	RuntimeVersion                      string
	LastMileSecurity                    bool
	ForwardSslSession                   bool
	UpdateStrategy                      string
	Replicas                            int
	GroupId                             string
	ArtifactId                          string
	ArtifactVersion                     string
	Packaging                           string
	Assets                              []interface{}
	DesiredState                        string
}
