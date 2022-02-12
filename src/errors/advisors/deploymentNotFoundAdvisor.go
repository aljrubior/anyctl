package advisors

import (
	"fmt"
	"github.com/aljrubior/anyctl/managers/entities"
	"strings"
)

func NewDeploymentNotFoundAdvisor(message string, deploymentName string, options *[]entities.DeploymentItemEntity) *DeploymentNotFoundAdvisor {
	return &DeploymentNotFoundAdvisor{
		message,
		deploymentName,
		options,
	}
}

type DeploymentNotFoundAdvisor struct {
	Message        string
	DeploymentName string
	Options        *[]entities.DeploymentItemEntity
}

func (this *DeploymentNotFoundAdvisor) Advise() {

	message := fmt.Sprintf("Error: %s", this.Message)
	advice := this.getAdvice(3)

	if advice == "" {
		fmt.Printf("%s.", message)
	} else {
		fmt.Printf("%s. %s", message, advice)
	}

	fmt.Printf("\n")
}

func (this *DeploymentNotFoundAdvisor) getAdvice(maxAdvice int) string {

	options := ""
	count := 0

	if this.Options == nil {
		return ""
	}

	for _, k := range *this.Options {

		if strings.Contains(k.Name, this.DeploymentName) {
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
