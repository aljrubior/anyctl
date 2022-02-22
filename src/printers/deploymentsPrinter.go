package printers

import (
	"fmt"
	"github.com/aljrubior/anyctl/managers/entities"
	"os"
	"text/tabwriter"
)

func NewDeploymentsPrinter(entities *[]entities.DeploymentItemEntity, targets *[]entities.TargetEntity) *DeploymentsPrinter {

	return &DeploymentsPrinter{
		entities: entities,
		targets:  targets,
	}
}

type DeploymentsPrinter struct {
	entities *[]entities.DeploymentItemEntity
	targets  *[]entities.TargetEntity
}

func (this *DeploymentsPrinter) Print() {

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	targetsMap := this.transformTargetEntitiesToMap(*this.targets)

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s", "NAME", "DSTATUS", "ASTATUS", "TARGET")

	for _, v := range *this.entities {
		fmt.Fprintf(w, "\n %s\t%s\t%s\t%s",
			v.Name,
			v.Status,
			v.Application.Status,
			targetsMap[v.Target.TargetId].GetName())
	}

	fmt.Fprintf(w, "\n")
}

func (this *DeploymentsPrinter) transformTargetEntitiesToMap(fromTargets []entities.TargetEntity) map[string]entities.TargetEntity {

	result := make(map[string]entities.TargetEntity)

	for _, v := range fromTargets {
		result[v.GetId()] = v
	}

	return result
}
