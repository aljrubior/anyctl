package printers

import (
	"fmt"
	"github.com/aljrubior/anyctl/managers/entities"
	"os"
	"text/tabwriter"
)

func NewOrganizationPrivateSpaceFabricsPrinter(entities *[]entities.OrganizationPrivateSpaceFabricEntity) *OrganizationPrivateSpaceFabricsPrinter {
	return &OrganizationPrivateSpaceFabricsPrinter{
		entities: entities,
	}
}

type OrganizationPrivateSpaceFabricsPrinter struct {
	entities *[]entities.OrganizationPrivateSpaceFabricEntity
}

func (this *OrganizationPrivateSpaceFabricsPrinter) Print() {
	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s", "NAME", "REGION", "STATUS", "STATUS MESSAGE")

	for _, v := range *this.entities {
		fmt.Fprintf(w, "\n %s\t%s\t%s\t%s",
			v.Name,
			v.Region,
			v.Status,
			v.StatusMessage,
		)
	}

	fmt.Fprintf(w, "\n")
}
