package manifests

import (
	"github.com/aljrubior/anyctl/clients/targets/response"
	"github.com/aljrubior/anyctl/model"
)

func NewRuntimeFabricTargetManifest(response response.RuntimeFabricTargetResponse) *RuntimeFabricTargetManifest {
	return &RuntimeFabricTargetManifest{
		ApiVersion: "v1",
		Kind:       "RuntimeFabricTarget",
		Metadata: model.Metadata{
			Name: response.Name,
		},
		Spec: response,
	}
}

type RuntimeFabricTargetManifest struct {
	ApiVersion string                               `yaml:"apiVersion"`
	Kind       string                               `yaml:"kind"`
	Metadata   model.Metadata                       `yaml:"metadata"`
	Spec       response.RuntimeFabricTargetResponse `yaml:"spec"`
}
