package requests

type Scheduler struct {
	FlowName   string `json:"flowName"`
	Enabled    bool   `json:"enabled"`
	Type       string `json:"type,omitempty"`
	StartDelay int    `json:"startDelay,omitempty"`
	Frequency  int    `json:"frequency,omitempty"`
	TimeUnit   string `json:"timeUnit,omitempty"`
	Expression string `json:"expression,omitempty"`
	TimeZone   string `json:"timeZone,omitempty"`
}
