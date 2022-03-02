package printers

import (
	"fmt"
	"github.com/aljrubior/anyctl/managers/entities"
	"os"
	"text/tabwriter"
)

func NewFabricsPrinter(entities *[]entities.FabricEntity, organizations map[string]*entities.OrganizationEntity) *FabricsPrinter {
	return &FabricsPrinter{
		entities:      entities,
		organizations: organizations,
	}
}

type FabricsPrinter struct {
	entities      *[]entities.FabricEntity
	organizations map[string]*entities.OrganizationEntity
}

func (this *FabricsPrinter) Print() {

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s", "NAME", "ORG NAME", "ORG TYPE", "VERSION", "REGION", "STATUS")

	for _, v := range *this.entities {

		fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s",
			v.Name,
			this.organizations[v.OrganizationId].Name,
			this.organizations[v.OrganizationId].Subscription.Type,
			v.Version,
			v.Region,
			v.Status)
	}

	fmt.Fprintf(w, "\n")
}
