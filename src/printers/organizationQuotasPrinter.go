package printers

import (
	"fmt"
	"github.com/aljrubior/anyctl/managers/entities"
	"os"
	"text/tabwriter"
)

func NewOrganizationQuotasPrinter(entity *entities.OrganizationEntity) *OrganizationQuotasPrinter {
	return &OrganizationQuotasPrinter{
		entity: entity,
	}

}

type OrganizationQuotasPrinter struct {
	entity *entities.OrganizationEntity
}

func (this *OrganizationQuotasPrinter) Print() {

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s", "NAME", "ASSIGNED", "USAGE")

	fmt.Fprintf(w, fmt.Sprintf("\n %s\t%s\t%s", "Production vCores", this.formatFloat(this.entity.Entitlements.VCoresProduction.Assigned), this.formatFloat(this.entity.Entitlements.VCoresProduction.Reassigned)))
	fmt.Fprintf(w, fmt.Sprintf("\n %s\t%s\t%s", "Sandbox vCores", this.formatFloat(this.entity.Entitlements.VCoresSandbox.Assigned), this.formatFloat(this.entity.Entitlements.VCoresSandbox.Reassigned)))
	fmt.Fprintf(w, fmt.Sprintf("\n %s\t%s\t%s", "Design vCores", this.formatFloat(this.entity.Entitlements.VCoresDesign.Assigned), this.formatFloat(this.entity.Entitlements.VCoresDesign.Reassigned)))
	fmt.Fprintf(w, fmt.Sprintf("\n %s\t%s\t%s", "Static IPs", this.formatFloat(this.entity.Entitlements.StaticIps.Assigned), this.formatFloat(this.entity.Entitlements.StaticIps.Reassigned)))
	fmt.Fprintf(w, fmt.Sprintf("\n %s\t%s\t%s", "VPCs", this.formatFloat(this.entity.Entitlements.Vpcs.Assigned), this.formatFloat(this.entity.Entitlements.Vpcs.Reassigned)))
	fmt.Fprintf(w, fmt.Sprintf("\n %s\t%s\t%s", "VPNs", this.formatFloat(this.entity.Entitlements.Vpns.Assigned), this.formatFloat(this.entity.Entitlements.Vpns.Reassigned)))
	fmt.Fprintf(w, fmt.Sprintf("\n %s\t%s\t%s", "Load Balancers", this.formatFloat(this.entity.Entitlements.LoadBalancer.Assigned), this.formatFloat(this.entity.Entitlements.LoadBalancer.Reassigned)))

	fmt.Fprintf(w, "\n")
}

func (this *OrganizationQuotasPrinter) formatFloat(value float64) string {

	if value > 0 {
		return fmt.Sprintf("%v", value)
	}

	return "-"
}
