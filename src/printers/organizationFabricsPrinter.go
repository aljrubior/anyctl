package printers

import (
	"fmt"
	"github.com/aljrubior/anyctl/managers/entities"
	"os"
	"text/tabwriter"
)

func NewOrganizationFabricsPrinter(entitites *[]entities.OrganizationFabricEntity) *OrganizationFabricsPrinter {

	return &OrganizationFabricsPrinter{
		entities: entitites,
	}
}

type OrganizationFabricsPrinter struct {
	entities *[]entities.OrganizationFabricEntity
}

func (this *OrganizationFabricsPrinter) Print() {

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s\t%s", "NAME", "REGION", "FABRIC VERSION", "STATUS", "AVAILABLE UPGRADE", "LEVEL", "DISTRIBUTION")

	for _, v := range *this.entities {
		fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s\t%s",
			v.Name,
			v.Region,
			v.Version,
			v.Status,
			v.AvailableUpgradeVersion,
			v.ClusterConfigurationLevel,
			this.getRuntimeFabricDistribution(&v),
		)
	}

	fmt.Fprintf(w, "\n")
}

func (this *OrganizationFabricsPrinter) getRuntimeFabricDistribution(entity *entities.OrganizationFabricEntity) string {
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
