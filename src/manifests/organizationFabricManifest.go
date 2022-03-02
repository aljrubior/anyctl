package manifests

import (
	"github.com/aljrubior/anyctl/clients/organizationRuntimeFabrics/response"
	"github.com/aljrubior/anyctl/model"
)

func NewOrganizationFabricManifest(response response.OrganizationFabricResponse) *OrganizationFabricManifest {
	return &OrganizationFabricManifest{
		ApiVersion: "v1",
		Kind:       "RuntimeFabric",
		Metadata: model.Metadata{
			Name: response.Name,
		},
		Spec: response,
	}
}

type OrganizationFabricManifest struct {
	ApiVersion string                              `json:"apiVersion"`
	Kind       string                              `json:"kind"`
	Metadata   model.Metadata                      `json:"metadata"`
	Spec       response.OrganizationFabricResponse `json:"spec"`
}
