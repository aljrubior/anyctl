package printers

import (
	"fmt"
	"github.com/aljrubior/anyctl/builders"
	"github.com/aljrubior/anyctl/managers/entities"
	"os"
	"text/tabwriter"
)

func NewOrganizationFabricPrinter(entity *entities.OrganizationFabricEntity) *OrganizationFabricPrinter {

	return &OrganizationFabricPrinter{
		entity: entity,
	}
}

type OrganizationFabricPrinter struct {
	entity *entities.OrganizationFabricEntity
}

func (this *OrganizationFabricPrinter) Print() {

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	nodeSummary := builders.NewRuntimeFabricNodeSummaryBuilder(&this.entity.Nodes).Build()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s", "NAME", "READY", "HEALTHY", "SCHEDULABLE", "CAPACITY", "REGION", "VERSION", "STATUS", "DISTRIBUTION")

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s",
		this.entity.Name,
		nodeSummary.Ready,
		nodeSummary.Healty,
		nodeSummary.Schedulable,
		nodeSummary.Capacity,
		this.entity.Region,
		this.entity.Version,
		this.entity.Status,
		this.getRuntimeFabricDistribution(this.entity),
	)

	fmt.Fprintf(w, "\n")
}

func (this *OrganizationFabricPrinter) getRuntimeFabricDistribution(entity *entities.OrganizationFabricEntity) string {
	switch entity.Vendor {
	case "aks":
		return "AKS"
	case "eks":
		return "EKS"
	case "gke":
		return "GKE"
	case "gravitational":
		return "APPLIANCE"
	case "rtfc":
		return "RTFC"
	default:
		return "Unknown"
	}
}
