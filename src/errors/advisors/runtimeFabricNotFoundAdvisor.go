package advisors

import (
	"fmt"
	"github.com/aljrubior/anyctl/managers/entities"
	"strings"
)

func NewOrganizationFabricNotFoundAdvisor(message string, runtimeFabricName string, options *[]entities.OrganizationFabricEntity) *OrganizationFabricNotFoundAdvisor {
	return &OrganizationFabricNotFoundAdvisor{
		Message:           message,
		RuntimeFabricName: runtimeFabricName,
		Options:           options,
	}
}

type OrganizationFabricNotFoundAdvisor struct {
	Message           string
	RuntimeFabricName string
	Options           *[]entities.OrganizationFabricEntity
}

func (this *OrganizationFabricNotFoundAdvisor) Advise() {

	message := fmt.Sprintf("Error: %s", this.Message)
	advice := this.getAdvice(3)

	if advice == "" {
		fmt.Println(message)
	} else {
		fmt.Printf("%s. %s", message, advice)
	}
}

func (this *OrganizationFabricNotFoundAdvisor) getAdvice(maxAdvice int) string {

	options := ""
	count := 0
	lastValue := ""

	for _, k := range *this.Options {
		if lastValue != k.Name &&
			strings.Contains(k.Name, this.RuntimeFabricName) {
			options += fmt.Sprintf("%s ", k.Name)
			lastValue = k.Name
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
