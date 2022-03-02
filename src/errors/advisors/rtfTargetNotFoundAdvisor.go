package advisors

import (
	"fmt"
	"github.com/aljrubior/anyctl/managers/entities"
	"strings"
)

func NewRtfTargetNotFoundAdvisor(message string, targetName string, options *[]entities.RtfTargetEntity) *RtfTargetNotFoundAdvisor {
	return &RtfTargetNotFoundAdvisor{
		message,
		targetName,
		options,
	}
}

type RtfTargetNotFoundAdvisor struct {
	Message    string
	TargetName string
	Options    *[]entities.RtfTargetEntity
}

func (this *RtfTargetNotFoundAdvisor) Advise() {
	message := fmt.Sprintf("ERROR: %s", this.Message)
	advice := this.getAdvice(3)

	if advice == "" {
		fmt.Println(message)
	} else {
		fmt.Printf("%s. %s", message, advice)
	}
}

func (this *RtfTargetNotFoundAdvisor) getAdvice(maxAdvice int) string {

	options := ""
	count := 0

	for _, k := range *this.Options {

		if strings.Contains(k.Name, this.TargetName) {
			options += fmt.Sprintf("%s ", k.Name)
			count++
		}

		if count >= maxAdvice {
			break
		}
	}

	return fmt.Sprintf("Maybe you mean [%s]\n", strings.Trim(options, " "))
}
