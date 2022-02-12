package requests

func NewDeploymentStartRequestBuilder() *DeploymentStartRequestBuilder {
	return &DeploymentStartRequestBuilder{}
}

type DeploymentStartRequestBuilder struct {
}

func (this DeploymentStartRequestBuilder) Build() *DeploymentRequest {
	return &DeploymentRequest{
		Application: &Application{
			DesiredState: "STARTED",
		},
	}
}
