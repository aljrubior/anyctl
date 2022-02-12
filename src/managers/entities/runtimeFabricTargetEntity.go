package entities

import "github.com/aljrubior/anyctl/clients/targets/response"

type RuntimeFabricTargetEntity struct {
	response.RuntimeFabricTargetResponse
}

func (this *RuntimeFabricTargetEntity) Id() string {
	return this.RuntimeFabricTargetResponse.GetId()
}

func (this *RuntimeFabricTargetEntity) Name() string {
	return this.RuntimeFabricTargetResponse.GetName()
}

func (this *RuntimeFabricTargetEntity) Type() string {
	return this.RuntimeFabricTargetResponse.GetType()
}
