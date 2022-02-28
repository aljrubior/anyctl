package requests

func NewSchedulerEnableRequestBuilder() *SchedulerEnableRequestBuilder {
	return &SchedulerEnableRequestBuilder{}
}

type SchedulerEnableRequestBuilder struct {
	schedulers []Scheduler
}

func (this *SchedulerEnableRequestBuilder) AddScheduler(flowName string, enabled bool) *SchedulerEnableRequestBuilder {

	this.schedulers = append(this.schedulers, Scheduler{
		FlowName: flowName,
		Enabled:  enabled,
	})

	return this
}

func (this SchedulerEnableRequestBuilder) Build() *SchedulerEnableRequest {

	return &SchedulerEnableRequest{
		Schedulers: this.schedulers,
	}
}
