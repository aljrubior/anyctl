package printers

import (
	"fmt"
	"github.com/aljrubior/anyctl/managers/entities"
	"os"
	"text/tabwriter"
)

func NewPrivateSpacePrinter(entity *entities.PrivateSpaceEntity) *PrivateSpacePrinter {

	return &PrivateSpacePrinter{
		entity: entity,
	}
}

type PrivateSpacePrinter struct {
	entity *entities.PrivateSpaceEntity
}

func (this *PrivateSpacePrinter) PrintManagedFirewallRules() {
	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s", "CIDR BLOCK", "PROTOCOL", "FROM PORT", "TO PORT", "TYPE")

	for _, v := range this.entity.ManagedFirewallRules {
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

func (this *PrivateSpacePrinter) Print() {

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s", "NAME", "REGION", "STATUS", "VERSION", "FLAVOR", "ENVIRONMENT TYPE")

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%v\t%v",
		this.entity.Name,
		this.entity.Region,
		this.entity.Status,
		this.entity.Version,
		this.entity.Flavor,
		this.entity.Environments.Type,
	)

	fmt.Fprintf(w, "\n")
}
