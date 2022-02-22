package printers

import (
	"fmt"
	"github.com/aljrubior/anyctl/managers/entities"
	"os"
	"strings"
	"text/tabwriter"
)

func NewOrganizationPrivateSpacePrinter(entity *entities.OrganizationPrivateSpaceEntity) *OrganizationPrivateSpacePrinter {
	return &OrganizationPrivateSpacePrinter{
		entity: entity,
	}
}

type OrganizationPrivateSpacePrinter struct {
	entity *entities.OrganizationPrivateSpaceEntity
}

func (this *OrganizationPrivateSpacePrinter) PrintNetwork() {

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s", "REGION", "CIDR BLOCK", "INBOUND IPS", "OUTBOUND IPS", "DNS")

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s",
		this.entity.Network.Region,
		this.entity.Network.CidrBlock,
		strings.Join(this.entity.Network.InboundStaticIps, ", "),
		strings.Join(this.entity.Network.OutboundStaticIps, ", "),
		this.entity.Network.DnsTarget,
	)

	fmt.Fprintf(w, "\n")
}

func (this *OrganizationPrivateSpacePrinter) PrintFirewallRules() {
	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s", "CIDR BLOCK", "PROTOCOL", "FROM PORT", "TO PORT", "TYPE")

	for _, v := range this.entity.FirewallRules {
		fmt.Fprintf(w, "\n %s\t%s\t%d\t%d\t%s",
			v.CidrBlock,
			v.Protocol,
			v.FromPort,
			v.ToPort,
			v.Type,
		)
	}

	fmt.Fprintf(w, "\n")
}
