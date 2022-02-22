package utils

import (
	"fmt"
	"github.com/aljrubior/anyctl/managers/entities"
	"os"
	"text/tabwriter"
)

func PrintOrganizationPrivateSpaces(privateSpaces *[]entities.OrganizationPrivateSpaceEntity) {

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s", "NAME", "REGION", "STATUS", "PROVISIONING STATUS")

	for _, v := range *privateSpaces {
		fmt.Fprintf(w, "\n %s\t%s\t%s\t%s",
			v.Name,
			v.Region,
			v.Status,
			v.Provisioning.Status,
		)
	}

	fmt.Fprintf(w, "\n")
}

func PrintOrganizationPrivateSpaceFabrics(privateSpaces *[]entities.OrganizationPrivateSpaceFabricEntity) {

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s", "NAME", "REGION", "STATUS", "STATUS MESSAGE")

	for _, v := range *privateSpaces {
		fmt.Fprintf(w, "\n %s\t%s\t%s\t%s",
			v.Name,
			v.Region,
			v.Status,
			v.StatusMessage,
		)
	}

	fmt.Fprintf(w, "\n")
}
