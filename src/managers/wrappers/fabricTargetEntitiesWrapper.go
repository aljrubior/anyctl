package wrappers

import (
	"fmt"
	"github.com/aljrubior/anyctl/clients/organizationRuntimeFabrics/response"
	"github.com/aljrubior/anyctl/managers/entities"
)

func NewFabricTargetEntitiesWrapper(targets *[]entities.FabricTargetEntity) *FabricTargetEntitiesWrapper {
	return &FabricTargetEntitiesWrapper{
		targets,
	}
}

type FabricTargetEntitiesWrapper struct {
	targets *[]entities.FabricTargetEntity
}

func (this *FabricTargetEntitiesWrapper) ExistsTarget(byName string) bool {

	for _, v := range *this.targets {
		if v.Name == byName {
			return true
		}
	}

	return false
}

func (this *FabricTargetEntitiesWrapper) GetTargetByName(name string) *entities.FabricTargetEntity {

	for _, v := range *this.targets {
		if v.Name == name {
			return &v
		}
	}

	return nil
}

func (this *FabricTargetEntitiesWrapper) GetLatestRuntimeVersion(forTargetName string) string {

	target := this.GetTargetByName(forTargetName)

	if target == nil {
		return ""
	}

	return this.GetLatestRuntimeVersionByType(&target.Runtimes, "mule")

}

func (this *FabricTargetEntitiesWrapper) GetLatestRuntimeVersionByType(runtimes *[]response.Runtime, byType string) string {

	mapOfRuntimes := this.mapRuntimesByType(runtimes)

	runtime := mapOfRuntimes[byType]

	if &runtime == nil {
		return ""
	}

	version := runtime.Versions[0]

	if &version == nil {
		return ""
	}

	return fmt.Sprintf("%s:%s", version.BaseVersion, version.Tag)
}

func (this *FabricTargetEntitiesWrapper) mapRuntimesByType(runtimes *[]response.Runtime) map[string]response.Runtime {

	result := make(map[string]response.Runtime)

	for _, v := range *runtimes {
		result[v.Type] = v
	}

	return result
}

func (this *FabricTargetEntitiesWrapper) mapRuntimeVersionsByBaseVersion(versions *[]response.RuntimeVersion) map[string]response.RuntimeVersion {
	result := make(map[string]response.RuntimeVersion)

	for _, v := range *versions {
		result[v.BaseVersion] = v
	}

	return result
}
