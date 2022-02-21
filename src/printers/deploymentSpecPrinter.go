package printers

import (
	"fmt"
	"github.com/aljrubior/anyctl/managers/entities"
	"os"
	"text/tabwriter"
	"time"
)

func NewDeploymentSpecPrinter(specs *[]entities.DeploymentSpecEntity) *DeploymentSpecPrinter {

	return &DeploymentSpecPrinter{
		specs: specs,
	}
}

type DeploymentSpecPrinter struct {
	specs *[]entities.DeploymentSpecEntity
}

func (this *DeploymentSpecPrinter) Print() {
	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s", "DATE", "CHANGES")

	for i, v := range *this.specs {
		createdAt := time.UnixMilli(v.CreatedAt).Format("2006-01-02T15:04:05")
		version := v.Version[:6]
		if i == 0 {
			version = fmt.Sprintf("%s (Last successful)", version)
		}

		fmt.Fprintf(w, fmt.Sprintf("\n %s\t%s",
			createdAt,
			version))
	}

	fmt.Fprintf(w, "\n")
}
