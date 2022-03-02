package printers

import (
	"fmt"
	"github.com/aljrubior/anyctl/managers/entities"
	"os"
	"text/tabwriter"
)

func NewOrganizationUsagePrinter(entities []*entities.OrganizationEntity) *OrganizationUsagePrinter {

	return &OrganizationUsagePrinter{
		entities: entities,
	}
}

type OrganizationUsagePrinter struct {
	entities []*entities.OrganizationEntity
}

func (this *OrganizationUsagePrinter) Print() {

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s",
		"ORG ID",
		"ORG NAME",
		"PROD VCORES",
		"SAND VCORES",
		"DESIGN VCORES",
		"STATIC IPS",
		"VPCS",
		"VPNS",
		"LB")

	for _, v := range this.entities {

		if this.shouldPrintOrgUsage(v) {
			fmt.Fprintf(w, fmt.Sprintf("\n %s\t%s\t%v\t%v\t%v\t%v\t%v\t%v\t%v",
				v.Id,
				v.Name,
				this.formatFloat(v.Entitlements.VCoresProduction.Reassigned),
				this.formatFloat(v.Entitlements.VCoresSandbox.Reassigned),
				this.formatFloat(v.Entitlements.VCoresDesign.Reassigned),
				this.formatFloat(v.Entitlements.StaticIps.Reassigned),
				this.formatFloat(v.Entitlements.Vpcs.Reassigned),
				this.formatFloat(v.Entitlements.Vpns.Reassigned),
				this.formatFloat(v.Entitlements.LoadBalancer.Reassigned)))
		}
	}

	fmt.Fprintf(w, "\n")
}

func (this *OrganizationUsagePrinter) formatFloat(value float64) string {

	if value > 0 {
		return fmt.Sprintf("%v", value)
	}

	return "-"
}

func (this *OrganizationUsagePrinter) shouldPrintOrgUsage(org *entities.OrganizationEntity) bool {

	if org.OrganizationResponse.Entitlements.VCoresProduction.Reassigned > 0.0 ||
		org.OrganizationResponse.Entitlements.VCoresSandbox.Reassigned > 0.0 ||
		org.OrganizationResponse.Entitlements.VCoresDesign.Reassigned > 0.0 ||
		org.OrganizationResponse.Entitlements.StaticIps.Reassigned > 0.0 ||
		org.OrganizationResponse.Entitlements.Vpcs.Reassigned > 0.0 ||
		org.OrganizationResponse.Entitlements.Vpns.Reassigned > 0.0 ||
		org.OrganizationResponse.Entitlements.LoadBalancer.Reassigned > 0.0 {
		return true
	}

	return false
}
