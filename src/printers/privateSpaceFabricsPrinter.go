package printers

import (
	"fmt"
	"github.com/aljrubior/anyctl/managers/entities"
	"os"
	"text/tabwriter"
)

func NewPrivateSpaceFabricsPrinter(entities *[]entities.FabricEntity) *PrivateSpaceFabricsPrinter {
	return &PrivateSpaceFabricsPrinter{
		entities: entities,
	}
}

type PrivateSpaceFabricsPrinter struct {
	entities *[]entities.FabricEntity
}

func (this *PrivateSpaceFabricsPrinter) Print() {
	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s\t%s", "NAME", "REGION", "VERSION", "STATUS", "LEVEL", "INFRA VERSION", "INFRA ID")

	for _, v := range *this.entities {
		fmt.Fprintf(w, "\n  %s\t%s\t%s\t%s\t%s\t%s\t%s",
			v.Name,
			v.Region,
			v.Version,
			v.Status,
			v.ClusterConfigurationLevel,
			v.InfraVersion,
			v.InfraDeploymentId)
	}

	fmt.Fprintf(w, "\n")
}
