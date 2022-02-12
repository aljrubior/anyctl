package requests

func NewSchedulerEnableRequestBuilder(flowName string, enabled bool) *SchedulerEnableRequestBuilder {
	return &SchedulerEnableRequestBuilder{
		flowName,
		enabled,
	}
}

type SchedulerEnableRequestBuilder struct {
	flowName string
	enabled  bool
}

func (this SchedulerEnableRequestBuilder) Build() *SchedulerEnableRequest {

	var request []Scheduler

	request = append(request, Scheduler{
		FlowName: this.flowName,
		Enabled:  this.enabled,
	})

	return &SchedulerEnableRequest{
		Schedulers: request,
	}
}
