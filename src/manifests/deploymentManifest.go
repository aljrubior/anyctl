package manifests

import (
	"github.com/aljrubior/anyctl/clients/deployments/response"
	"github.com/aljrubior/anyctl/model"
)

func NewDeploymentManifest(response response.DeploymentResponse) *DeploymentManifest {
	return &DeploymentManifest{
		ApiVersion: "v1",
		Kind:       "Deployment",
		Metadata: model.Metadata{
			Name: response.Name,
		},
		Spec: response,
	}
}

type DeploymentManifest struct {
	ApiVersion string                      `yaml:"apiVersion"`
	Kind       string                      `yaml:"kind"`
	Metadata   model.Metadata              `yaml:"metadata"`
	Spec       response.DeploymentResponse `yaml:"spec"`
}
