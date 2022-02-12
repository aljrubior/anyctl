package requests

func NewDeploymentStopRequestBuilder() *DeploymentStopRequestBuilder {
	return &DeploymentStopRequestBuilder{}
}

type DeploymentStopRequestBuilder struct {
}

func (this DeploymentStopRequestBuilder) Build() *DeploymentRequest {
	return &DeploymentRequest{
		Application: &Application{
			DesiredState: "STOPPED",
		},
	}
}
