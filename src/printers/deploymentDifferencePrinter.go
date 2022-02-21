package printers

import (
	"fmt"
	"github.com/aljrubior/anyctl/comparators"
	"os"
	"text/tabwriter"
)

func NewDeploymentDifferencePrinter(differences []comparators.DeploymentDifference) *DeploymentDifferencePrinter {
	return &DeploymentDifferencePrinter{
		differences: differences,
	}
}

type DeploymentDifferencePrinter struct {
	differences []comparators.DeploymentDifference
}

func (this *DeploymentDifferencePrinter) Print() {
	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	var addCount, deleteCount, changeCount int

	for _, v := range this.differences {
		value := ""
		keyName := fmt.Sprintf("%v:", v.KeyName)

		switch v.Operator {
		case "~":
			if v.LeftValue == nil && v.RightValue == nil {
				value = ""
			} else {
				value = fmt.Sprintf("%s --> %v", v.LeftValue, v.RightValue)
				changeCount++
			}
		case "-":
			value = fmt.Sprintf("%v", v.RightValue)
			deleteCount++
		case "+":
			value = fmt.Sprintf("%v", v.RightValue)
			addCount++
		}

		fmt.Fprintf(w, "%s %s %s %s\n", this.spaces(v.Depth), v.Operator, keyName, value)
	}

	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "Result: %v to add, %v to change, %v to delete.", addCount, changeCount, deleteCount)
	fmt.Fprintf(w, "\n")
}

func (this *DeploymentDifferencePrinter) spaces(depth int) string {
	spaces := ""

	for i := 0; i <= depth; i++ {
		spaces += "  "
	}

	return spaces
}
