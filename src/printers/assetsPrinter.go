package printers

import (
	"fmt"
	"github.com/aljrubior/anyctl/managers/entities"
	"os"
	"text/tabwriter"
)

func NewAssetsPrinter(entities *[]entities.AssetEntity) *AssetsPrinter {

	return &AssetsPrinter{
		entities: entities,
	}
}

type AssetsPrinter struct {
	entities *[]entities.AssetEntity
}

func (this *AssetsPrinter) Print() {

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s", "NAME", "VERSION", "STATUS", "RUNTIME", "UPDATED AT", "CREATED BY")

	for _, v := range *this.entities {
		fmt.Fprintf(w, fmt.Sprintf("\n %s\t%s\t%s\t%s\t%s\t%s",
			v.Name,
			v.Version,
			v.Status,
			v.RuntimeVersion,
			v.UpdatedAt,
			v.CreatedBy.UserName))
	}

	fmt.Fprintf(w, "\n")
}
