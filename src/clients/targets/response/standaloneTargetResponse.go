package response

import "strconv"

type StandaloneTargetResponse struct {
	Id      int               `json:"id"`
	Name    string            `json:"name"`
	Kind    string            `json:"type"`
	Details StandaloneDetails `json:"details"`
	Status  string            `json:"status"`
}

func (this *StandaloneTargetResponse) GetId() string {
	return strconv.Itoa(this.Id)
}

func (this *StandaloneTargetResponse) GetName() string {
	return this.Name
}

func (this *StandaloneTargetResponse) GetType() string {
	return this.Kind
}
