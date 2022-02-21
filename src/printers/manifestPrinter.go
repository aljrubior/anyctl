package printers

import (
	"encoding/json"
	"github.com/aljrubior/anyctl/manifests"
	"sigs.k8s.io/yaml"
)

func NewDeploymentManifestPrinter(manifest *manifests.DeploymentManifest) (*ManifestPrinter, error) {

	dataAsJson, err := json.Marshal(*manifest)

	if err != nil {
		return nil, err
	}

	dataAsYaml, err := yaml.JSONToYAML(dataAsJson)

	return &ManifestPrinter{
		data: dataAsYaml,
	}, nil

}

func NewRuntimeFabricTargetManifestPrinter(manifest *manifests.RuntimeFabricTargetManifest) (*ManifestPrinter, error) {

	dataAsJson, err := json.Marshal(*manifest)

	if err != nil {
		return nil, err
	}

	dataAsYaml, err := yaml.JSONToYAML(dataAsJson)

	return &ManifestPrinter{
		data: dataAsYaml,
	}, nil

}

func NewStandaloneTargetManifestPrinter(manifest *manifests.StandaloneTargetManifest) (*ManifestPrinter, error) {

	dataAsJson, err := json.Marshal(*manifest)

	if err != nil {
		return nil, err
	}

	dataAsYaml, err := yaml.JSONToYAML(dataAsJson)

	return &ManifestPrinter{
		data: dataAsYaml,
	}, nil

}

type ManifestPrinter struct {
	data []byte
}

func (this *ManifestPrinter) Print() {
	println(string(this.data))
}
