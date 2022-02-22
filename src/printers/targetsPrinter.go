package printers

import (
	"fmt"
	"github.com/aljrubior/anyctl/managers/entities"
	"os"
	"text/tabwriter"
)

func NewTargetsPrinter(entities *[]entities.TargetEntity) *TargetsPrinter {
	return &TargetsPrinter{
		entities: entities,
	}
}

type TargetsPrinter struct {
	entities *[]entities.TargetEntity
}

func (this *TargetsPrinter) Print() {

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s", "ID", "NAME", "TYPE")

	for _, v := range *this.entities {
		fmt.Fprintf(w, "\n %s\t%s\t%s",
			v.GetId(),
			v.GetName(),
			v.GetType())
	}

	fmt.Fprintf(w, "\n")
}
