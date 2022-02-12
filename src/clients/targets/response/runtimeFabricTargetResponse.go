package response

type RuntimeFabricTargetResponse struct {
	Id      string        `json:"id"`
	Name    string        `json:"name"`
	Kind    string        `json:"type"`
	Details []TagetDetail `json:"details"`
}

func (this *RuntimeFabricTargetResponse) GetId() string {
	return this.Id
}

func (this *RuntimeFabricTargetResponse) GetName() string {
	return this.Name
}

func (this *RuntimeFabricTargetResponse) GetType() string {
	return this.Kind
}
