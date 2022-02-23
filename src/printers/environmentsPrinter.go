package printers

import (
	"fmt"
	"github.com/aljrubior/anyctl/model"
	"os"
	"text/tabwriter"
)

func NewEnvironmentsPrinter(environments *[]model.Environment) *EnvironmentsPrinter {
	return &EnvironmentsPrinter{
		environments: environments,
	}
}

type EnvironmentsPrinter struct {
	environments *[]model.Environment
}

func (this *EnvironmentsPrinter) Print() {

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s", "ID", "NAME", "TYPE")

	for _, v := range *this.environments {
		fmt.Fprintf(w, "\n %s\t%s\t%s",
			v.Id,
			v.Name,
			v.Kind,
		)
	}

	fmt.Fprintf(w, "\n")
}
