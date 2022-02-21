package requests

import "github.com/aljrubior/anyctl/clients/deployments/response"

func NewDeploymentScaleRequestBuilder(desiredReplicas int, target response.Target) *DeploymentScaleRequestBuilder {
	return &DeploymentScaleRequestBuilder{
		desiredReplicas: desiredReplicas,
		target:          target,
	}
}

type DeploymentScaleRequestBuilder struct {
	target          response.Target
	desiredReplicas int
}

func (this DeploymentScaleRequestBuilder) Build() *DeploymentRequest {

	return &DeploymentRequest{
		Target: Target{
			Provider: this.target.Provider,
			TargetId: this.target.TargetId,
			Replicas: this.desiredReplicas,
			DeploymentSettings: DeploymentSettings{
				Resources: Resources{
					Cpu: ResourceItem{
						Limit:    this.target.DeploymentSettings.Resources.Cpu.Limit,
						Reserved: this.target.DeploymentSettings.Resources.Cpu.Reserved,
					},
					Memory: ResourceItem{
						Limit:    this.target.DeploymentSettings.Resources.Memory.Limit,
						Reserved: this.target.DeploymentSettings.Resources.Memory.Reserved,
					},
				},
				Clustered:                           this.target.DeploymentSettings.Clustered,
				EnforceDeployingReplicasAcrossNodes: this.target.DeploymentSettings.EnforceDeployingReplicasAcrossNodes,
				Http: Http{
					Inbound: Inbound{
						PublicUrl:   this.target.DeploymentSettings.Http.Inbound.PublicUrl,
						PathRewrite: this.target.DeploymentSettings.Http.Inbound.PathRewrite,
					},
				},
				Jvm:               Jvm{},
				RuntimeVersion:    this.target.DeploymentSettings.RuntimeVersion,
				LastMileSecurity:  this.target.DeploymentSettings.LastMileSecurity,
				ForwardSslSession: this.target.DeploymentSettings.ForwardSslSession,
				UpdateStrategy:    this.target.DeploymentSettings.UpdateStrategy,
			},
		},
	}
}
