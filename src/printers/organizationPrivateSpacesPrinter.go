package printers

import (
	"fmt"
	"github.com/aljrubior/anyctl/managers/entities"
	"os"
	"text/tabwriter"
)

func NewOrganizationPrivateSpacesPrinter(entities *[]entities.OrganizationPrivateSpaceEntity) *OrganizationPrivateSpacesPrinter {
	return &OrganizationPrivateSpacesPrinter{
		entities: entities,
	}
}

type OrganizationPrivateSpacesPrinter struct {
	entities *[]entities.OrganizationPrivateSpaceEntity
}

func (this *OrganizationPrivateSpacesPrinter) Print() {
	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s", "NAME", "REGION", "STATUS", "PROVISIONING STATUS")

	for _, v := range *this.entities {
		fmt.Fprintf(w, "\n %s\t%s\t%s\t%s",
			v.Name,
			v.Region,
			v.Status,
			v.Provisioning.Status,
		)
	}

	fmt.Fprintf(w, "\n")
}
