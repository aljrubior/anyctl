package manifests

import (
	"github.com/aljrubior/anyctl/clients/fabrics/response"
	"github.com/aljrubior/anyctl/model"
)

func NewFabricManifest(response response.FabricResponse) *FabricManifest {
	return &FabricManifest{
		ApiVersion: "v1",
		Kind:       "Fabric",
		Metadata: model.Metadata{
			Name: response.Name,
		},
		Spec: response,
	}
}

type FabricManifest struct {
	ApiVersion string                  `json:"apiVersion"`
	Kind       string                  `json:"kind"`
	Metadata   model.Metadata          `json:"metadata"`
	Spec       response.FabricResponse `json:"spec"`
}
