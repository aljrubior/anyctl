package conf

type SchedulerClientConfig struct {
	Protocol                 string
	Host                     string
	Port                     int
	SchedulersPathTemplate   string
	SchedulerPathTemplate    string
	RunSchedulerPathTemplate string
}
