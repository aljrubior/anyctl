package printers

import (
	"fmt"
	"github.com/aljrubior/anyctl/managers/entities"
	"os"
	"text/tabwriter"
)

func NewSharedSpacesPrinter(entities *[]entities.SharedSpaceEntity) *SharedSpacesPrinter {

	return &SharedSpacesPrinter{
		entities: entities,
	}
}

type SharedSpacesPrinter struct {
	entities *[]entities.SharedSpaceEntity
}

func (this *SharedSpacesPrinter) Print() {

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s", "NAME", "REGION", "FLAVOR", "STATUS")

	for _, v := range *this.entities {
		fmt.Fprintf(w, "\n %s\t%s\t%s\t%s",
			v.Name,
			v.Region,
			v.Flavor,
			v.Status,
		)
	}

	fmt.Fprintf(w, "\n")
}
