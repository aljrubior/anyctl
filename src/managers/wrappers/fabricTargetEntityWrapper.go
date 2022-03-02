package wrappers

import (
	"fmt"
	"github.com/aljrubior/anyctl/clients/organizationRuntimeFabrics/response"
	"github.com/aljrubior/anyctl/managers/entities"
	"github.com/pkg/errors"
)

func NewFabricTargetEntityWrapper(target *entities.FabricTargetEntity) *FabricTargetEntityWrapper {
	return &FabricTargetEntityWrapper{
		target,
	}
}

type FabricTargetEntityWrapper struct {
	target *entities.FabricTargetEntity
}

func (this *FabricTargetEntityWrapper) GetRuntimeVersionRef(forType, forBaseVersion string) string {

	if forBaseVersion == "" {
		return this.GetLatestRuntimeVersionRef(forType)
	}

	return this.GetLatestTagForRuntimeVersionRef(forType, forBaseVersion)

}

func (this *FabricTargetEntityWrapper) GetLatestRuntimeVersionRef(forType string) string {

	return this.GetLatestRuntimeVersionByType(forType)

}

func (this *FabricTargetEntityWrapper) GetLatestTagForRuntimeVersionRef(forType, runtimeVersion string) string {

	mapOfRuntimes := this.mapRuntimesByType(this.target.Runtimes)

	runtime, exist := mapOfRuntimes[forType]

	if !exist {
		errors.New(fmt.Sprintf("Runtime type not found '%s'", runtimeVersion))
	}

	mapOfVersions := this.mapRuntimeVersionsByBaseVersion(runtime.Versions)

	version, exist := mapOfVersions[runtimeVersion]

	if !exist {
		errors.New(fmt.Sprintf("Runtime version not found '%s'", runtimeVersion))
	}

	return fmt.Sprintf("%s:%s", version.BaseVersion, version.Tag)

}

func (this *FabricTargetEntityWrapper) GetLatestRuntimeVersionByType(forType string) string {

	mapOfRuntimes := this.mapRuntimesByType(this.target.Runtimes)

	runtime := mapOfRuntimes[forType]

	if &runtime == nil {
		return ""
	}

	version := runtime.Versions[0]

	if &version == nil {
		return ""
	}

	return fmt.Sprintf("%s:%s", version.BaseVersion, version.Tag)
}

func (this *FabricTargetEntityWrapper) mapRuntimesByType(runtimes []response.Runtime) map[string]response.Runtime {

	result := make(map[string]response.Runtime)

	for _, v := range runtimes {
		result[v.Type] = v
	}

	return result
}

func (this *FabricTargetEntityWrapper) mapRuntimeVersionsByBaseVersion(versions []response.RuntimeVersion) map[string]response.RuntimeVersion {
	result := make(map[string]response.RuntimeVersion)

	for _, v := range versions {
		result[v.BaseVersion] = v
	}

	return result
}
