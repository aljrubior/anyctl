package entities

import "github.com/aljrubior/anyctl/clients/schedulers/response"

func NewSchedulerEntitiesBuilder(response *[]response.Scheduler) *SchedulerEntitiesBuilder {
	return &SchedulerEntitiesBuilder{
		response: response,
	}
}

type SchedulerEntitiesBuilder struct {
	response *[]response.Scheduler
}

func (this SchedulerEntitiesBuilder) Build() *[]SchedulerEntity {

	var schedulers []SchedulerEntity

	for _, v := range *this.response {
		schedulers = append(schedulers, SchedulerEntity{
			v,
		})
	}

	return &schedulers
}
