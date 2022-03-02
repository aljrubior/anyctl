package advisors

import (
	"fmt"
	"github.com/aljrubior/anyctl/managers/entities"
	"strings"
)

func NewSchedulerNotFoundAdvisor(message string, flowName string, options *[]entities.SchedulerEntity) *SchedulerNotFoundAdvisor {
	return &SchedulerNotFoundAdvisor{
		Message:  message,
		FlowName: flowName,
		Options:  options,
	}
}

type SchedulerNotFoundAdvisor struct {
	Message  string
	FlowName string
	Options  *[]entities.SchedulerEntity
}

func (this *SchedulerNotFoundAdvisor) Advise() {

	message := fmt.Sprintf("ERROR: %s", this.Message)
	advice := this.getAdvice(3)

	if advice == "" {
		fmt.Println(message)
	} else {
		fmt.Printf("%s. %s", message, advice)
	}
}

func (this *SchedulerNotFoundAdvisor) getAdvice(maxAdvice int) string {

	options := ""
	count := 0

	for _, k := range *this.Options {
		if strings.Contains(k.FlowName, this.FlowName) {
			options += fmt.Sprintf("%s ", k.FlowName)
			count++
		}

		if count >= maxAdvice {
			break
		}
	}

	if options == "" {
		return options
	}

	return fmt.Sprintf("Maybe you mean [%s]\n", strings.Trim(options, " "))
}
