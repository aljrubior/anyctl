package manifests

import (
	"github.com/aljrubior/anyctl/clients/sharedspaces/response"
	"github.com/aljrubior/anyctl/model"
)

func NewSharedSpaceManifest(response response.SharedSpaceResponse) *SharedSpaceManifest {
	return &SharedSpaceManifest{
		ApiVersion: "v1",
		Kind:       "SharedSpace",
		Metadata: model.Metadata{
			Name: response.Name,
		},
		Spec: response,
	}
}

type SharedSpaceManifest struct {
	ApiVersion string                       `yaml:"apiVersion"`
	Kind       string                       `yaml:"kind"`
	Metadata   model.Metadata               `yaml:"metadata"`
	Spec       response.SharedSpaceResponse `yaml:"spec"`
}
