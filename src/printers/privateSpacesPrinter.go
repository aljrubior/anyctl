package printers

import (
	"fmt"
	"github.com/aljrubior/anyctl/managers/entities"
	"os"
	"text/tabwriter"
)

func NewPrivateSpacesPrinter(entities *[]entities.PrivateSpaceEntity) *PrivateSpacesPrinter {
	return &PrivateSpacesPrinter{
		entities: entities,
	}
}

type PrivateSpacesPrinter struct {
	entities *[]entities.PrivateSpaceEntity
}

func (this *PrivateSpacesPrinter) Print() {
	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s", "NAME", "REGION", "STATUS", "VERSION", "FLAVOR", "ENVIRONMENT TYPE")

	for _, v := range *this.entities {
		fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%v\t%v",
			v.Name,
			v.Region,
			v.Status,
			v.Version,
			v.Flavor,
			v.Environments.Type)
	}

	fmt.Fprintf(w, "\n")
}
