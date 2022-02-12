package manifests

import (
	"github.com/aljrubior/anyctl/clients/targets/response"
	"github.com/aljrubior/anyctl/model"
)

func NewStandaloneTargetManifest(response response.StandaloneTargetResponse) *StandaloneTargetManifest {
	return &StandaloneTargetManifest{
		ApiVersion: "v1",
		Kind:       "StandaloneTarget",
		Metadata: model.Metadata{
			Name: response.Name,
		},
		Spec: response,
	}
}

type StandaloneTargetManifest struct {
	ApiVersion string                            `yaml:"apiVersion"`
	Kind       string                            `yaml:"kind"`
	Metadata   model.Metadata                    `yaml:"metadata"`
	Spec       response.StandaloneTargetResponse `yaml:"spec"`
}
