package advisors

import (
	"fmt"
	"github.com/aljrubior/anyctl/managers/entities"
	"strings"
)

func NewAssetGroupNotFoundAdvisor(message, groupId, assetName string, options *[]entities.AssetEntity) *AssetGroupNotFoundAdvisor {
	return &AssetGroupNotFoundAdvisor{
		message,
		groupId,
		assetName,
		options,
	}
}

type AssetGroupNotFoundAdvisor struct {
	Message   string
	GroupId   string
	AssetName string
	Options   *[]entities.AssetEntity
}

func (this *AssetGroupNotFoundAdvisor) Advise() {

	message := fmt.Sprintf("Error: %s", this.Message)
	advice := this.getAdvice(3)

	if advice == "" {
		fmt.Println(message)
	} else {
		fmt.Printf("%s. %s", message, advice)
	}
}

func (this *AssetGroupNotFoundAdvisor) getAdvice(maxAdvice int) string {

	options := ""
	count := 0
	lastValue := ""

	for _, k := range *this.Options {
		if lastValue != k.Name &&
			strings.Contains(k.GroupId, this.GroupId) {
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
