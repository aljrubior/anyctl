package wrappers

import "github.com/aljrubior/anyctl/managers/entities"

func NewTargetEntityWrapper(targetEntity entities.TargetEntity) *TargetEntityWrapper {
	return &TargetEntityWrapper{
		targetEntity,
	}
}

type TargetEntityWrapper struct {
	targetEntity entities.TargetEntity
}

func (this *TargetEntityWrapper) IsStandaloneTargetEntity() bool {

	_, ok := this.targetEntity.(*entities.StandaloneTargetEntity)

	return ok
}

func (this *TargetEntityWrapper) IsRuntimeFabricTargetEntity() bool {

	_, ok := this.targetEntity.(*entities.RuntimeFabricTargetEntity)

	return ok
}

func (this *TargetEntityWrapper) GetRuntimeFabricTargetEntity() (*entities.RuntimeFabricTargetEntity, bool) {

	entity, ok := this.targetEntity.(*entities.RuntimeFabricTargetEntity)

	return entity, ok
}

func (this *TargetEntityWrapper) GetStandaloneTargetEntity() (*entities.StandaloneTargetEntity, bool) {

	entity, ok := this.targetEntity.(*entities.StandaloneTargetEntity)

	return entity, ok
}
