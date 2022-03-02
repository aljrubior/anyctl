package printers

import (
	"fmt"
	"github.com/aljrubior/anyctl/managers/entities"
	"os"
	"text/tabwriter"
)

func NewSharedSpacePrinter(entity *entities.SharedSpaceEntity, privateSpace *entities.PrivateSpaceEntity) *SharedSpacePrinter {

	return &SharedSpacePrinter{
		entity:       entity,
		privateSpace: privateSpace,
	}
}

type SharedSpacePrinter struct {
	entity       *entities.SharedSpaceEntity
	privateSpace *entities.PrivateSpaceEntity
}

func (this *SharedSpacePrinter) Print() {
	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s\t%s", "NAME", "REGION", "FLAVOR", "STATUS", "ADVERTISED", "REQUIRES PERMISSION", "PRIVATE SPACE")

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%v\t%v\t%s",
		this.entity.Name,
		this.entity.Region,
		this.entity.Flavor,
		this.entity.Status,
		this.entity.IsAdvertised,
		this.entity.RequiresPermission,
		this.privateSpace.Name,
	)

	fmt.Fprintf(w, "\n")
}
