package printers

import (
	"fmt"
	"github.com/aljrubior/anyctl/managers/entities"
	"github.com/aljrubior/anyctl/managers/wrappers"
	"os"
	"text/tabwriter"
)

func NewTargetPrinter(entity *entities.TargetEntity) *TargetPrinter {
	return &TargetPrinter{
		entity: entity,
	}
}

type TargetPrinter struct {
	entity *entities.TargetEntity
}

func (this *TargetPrinter) PrintTargetSupportedVersions() {

	wrapper := wrappers.NewTargetEntityWrapper(*this.entity)

	target, ok := wrapper.GetRuntimeFabricTargetEntity()

	if !ok {
		println("Supported versions is not available on this target type")
		return
	}

	if len(target.Details) == 1 && len(target.Details[0].SupportedRuntimes) == 0 {
		println("No runtime versions found in this target.")
		return
	}

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s", "BASE VERSION", "TAG", "MINIMUM TAG")

	for _, v := range target.Details {
		for _, x := range v.SupportedRuntimes {
			fmt.Fprintf(w, "\n %s\t%s\t%s",
				x.BaseVersion,
				x.Tag,
				x.MinimumTag)
		}
	}

	fmt.Fprintf(w, "\n")
}

func (this *TargetPrinter) PrintStandaloneDetails() {

	wrapper := wrappers.NewTargetEntityWrapper(*this.entity)

	target, ok := wrapper.GetStandaloneTargetEntity()

	if !ok {
		println("Details is not available on this target type")
		return
	}

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s", "RUNTIME VERSION", "TYPE", "AGENT VERSION", "STATUS")

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s",
		target.Details.RuntimeVersion,
		target.Details.Type,
		target.Details.AgentVersion,
		target.Status)

	fmt.Fprintf(w, "\n")
}
