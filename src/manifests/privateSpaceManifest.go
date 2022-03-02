package manifests

import (
	"github.com/aljrubior/anyctl/clients/privateSpaces/response"
	"github.com/aljrubior/anyctl/model"
)

func NewPrivateSpaceManifest(response response.PrivateSpaceResponse) *PrivateSpaceManifest {
	return &PrivateSpaceManifest{
		ApiVersion: "v1",
		Kind:       "PrivateSpace",
		Metadata: model.Metadata{
			Name: response.Name,
		},
		Spec: response,
	}
}

type PrivateSpaceManifest struct {
	ApiVersion string                        `yaml:"apiVersion"`
	Kind       string                        `yaml:"kind"`
	Metadata   model.Metadata                `yaml:"metadata"`
	Spec       response.PrivateSpaceResponse `yaml:"spec"`
}
