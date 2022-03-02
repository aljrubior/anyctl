package manifests

import (
	"github.com/aljrubior/anyctl/clients/organizationPrivateSpaces/response"
	"github.com/aljrubior/anyctl/model"
)

func NewOrganizationPrivateSpaceManifest(response response.OrganizationPrivateSpaceResponse) *OrganizationPrivateSpaceManifest {
	return &OrganizationPrivateSpaceManifest{
		ApiVersion: "v1",
		Kind:       "PrivateSpace",
		Metadata: model.Metadata{
			Name: response.Name,
		},
		Spec: response,
	}
}

type OrganizationPrivateSpaceManifest struct {
	ApiVersion string                                    `yaml:"apiVersion"`
	Kind       string                                    `yaml:"kind"`
	Metadata   model.Metadata                            `yaml:"metadata"`
	Spec       response.OrganizationPrivateSpaceResponse `yaml:"spec"`
}
