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
	ApiVersion string                      `json:"apiVersion"`
	Kind       string                      `json:"kind"`
	Metadata   model.Metadata              `json:"metadata"`
	Spec       response.DeploymentResponse `json:"spec"`
}
