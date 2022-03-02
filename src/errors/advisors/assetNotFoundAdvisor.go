package advisors

import (
	"fmt"
	"github.com/aljrubior/anyctl/managers/entities"
	"strings"
)

func NewAssetNotFoundAdvisor(message string, assetName string, options *[]entities.AssetEntity) *AssetNotFoundAdvisor {
	return &AssetNotFoundAdvisor{
		Message:   message,
		AssetName: assetName,
		Options:   options,
	}
}

type AssetNotFoundAdvisor struct {
	Message   string
	AssetName string
	Options   *[]entities.AssetEntity
}

func (this *AssetNotFoundAdvisor) Advise() {

	message := fmt.Sprintf("Error: %s", this.Message)
	advice := this.getAdvice(3)

	if advice == "" {
		fmt.Println(message)
	} else {
		fmt.Printf("%s. %s", message, advice)
	}
}

func (this *AssetNotFoundAdvisor) getAdvice(maxAdvice int) string {

	options := ""
	count := 0
	lastValue := ""

	if this.Options == nil {
		return fmt.Sprintf("Try with 'anyctl assets get %s'", this.getAssetName())
	}

	for _, k := range *this.Options {
		if lastValue != k.Name &&
			strings.Contains(k.Name, this.AssetName) {
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

func (this *AssetNotFoundAdvisor) getAssetName() string {

	asset := strings.Split(this.AssetName, ":")

	switch len(asset) {
	case 1, 2:
		return asset[0]
	case 3:
		return asset[1]
	default:
		return this.AssetName
	}
}
