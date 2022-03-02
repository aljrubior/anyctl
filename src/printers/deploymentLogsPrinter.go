package printers

import (
	"fmt"
	"github.com/aljrubior/anyctl/managers/entities"
	"os"
	"text/tabwriter"
	"time"
)

func NewDeploymentLogsPrinter(entities *[]entities.DeploymentLogMessageEntity) *DeploymentLogsPrinter {

	return &DeploymentLogsPrinter{
		entities: entities,
	}
}

type DeploymentLogsPrinter struct {
	entities *[]entities.DeploymentLogMessageEntity
}

func (this *DeploymentLogsPrinter) Print() {

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	for _, v := range *this.entities {

		date := time.UnixMilli(v.Timestamp).Format("2006-01-02T15:04:05")

		fmt.Fprintf(w, "\n %v\t%s\t%s\t%s",
			date,
			v.ReplicaId,
			v.LogLevel,
			v.Message,
		)
	}

	fmt.Fprintf(w, "\n")

}
