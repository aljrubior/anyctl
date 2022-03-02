package advisors

import (
	"fmt"
	"github.com/aljrubior/anyctl/model"
	"strings"
)

func NewAnypointEnvironmentNotFoundAdvisor(message string, environmentName string, options *[]model.Environment) *AnypointEnvironmentNotFoundAdvisor {
	return &AnypointEnvironmentNotFoundAdvisor{
		Message:         message,
		EnvironmentName: environmentName,
		Options:         options,
	}
}

type AnypointEnvironmentNotFoundAdvisor struct {
	Message         string
	EnvironmentName string
	Options         *[]model.Environment
}

func (this *AnypointEnvironmentNotFoundAdvisor) Advise() {

	message := fmt.Sprintf("Error: %s", this.Message)
	advice := this.getAdvice(3)

	if advice == "" {
		fmt.Printf("%s.", message)
	} else {
		fmt.Printf("%s. %s", message, advice)
	}
}

func (this *AnypointEnvironmentNotFoundAdvisor) getAdvice(maxAdvice int) string {

	options := ""
	count := 0

	for _, k := range *this.Options {
		if strings.Contains(k.Name, this.EnvironmentName) {
			options += fmt.Sprintf("%s ", k.Name)
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
