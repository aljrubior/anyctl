package printers

import (
	"fmt"
	"github.com/aljrubior/anyctl/managers/entities"
	"os"
	"text/tabwriter"
)

func NewFabricPrinter(entity *entities.FabricEntity) *FabricPrinter {

	return &FabricPrinter{
		entity: entity,
	}
}

type FabricPrinter struct {
	entity *entities.FabricEntity
}

func (this *FabricPrinter) PrintVersionInformation() {
	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s", "VERSION", "DESIRED", "UPGRADE AVAILABLE", "KUBERNETES", "INFRA", "DESIRED INFRA")

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s",
		this.entity.Version,
		this.entity.DesiredVersion,
		this.entity.AvailableUpgradeVersion,
		this.entity.KubernetesVersion,
		this.entity.InfraVersion,
		this.entity.DesiredInfraVersion)

	fmt.Fprintf(w, "\n")
}
