package advisors

import (
	"fmt"
	"github.com/aljrubior/anyctl/managers/entities"
	"strings"
)

func NewTargetNotFoundAdvisor(message string, targetName string, options *[]entities.TargetEntity) *TargetNotFoundAdvisor {
	return &TargetNotFoundAdvisor{
		message,
		targetName,
		options,
	}
}

type TargetNotFoundAdvisor struct {
	Message    string
	TargetName string
	Options    *[]entities.TargetEntity
}

func (this *TargetNotFoundAdvisor) Advise() {
	message := fmt.Sprintf("Error: %s", this.Message)
	advice := this.getAdvice(3)

	if advice == "" {
		fmt.Println(message)
	} else {
		fmt.Printf("%s. %s", message, advice)
	}
}

func (this *TargetNotFoundAdvisor) getAdvice(maxAdvice int) string {

	options := ""
	count := 0

	for _, k := range *this.Options {

		if strings.Contains(k.GetName(), this.TargetName) {
			options += fmt.Sprintf("%s ", k.GetName())
			count++
		}

		if count >= maxAdvice {
			break
		}
	}

	return fmt.Sprintf("Maybe you mean [%s]\n", strings.Trim(options, " "))
}
