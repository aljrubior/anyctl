package utils

import (
	"fmt"
	"github.com/aljrubior/anyctl/builders"
	"github.com/aljrubior/anyctl/clients/fabrics/response"
	"github.com/aljrubior/anyctl/managers/entities"
	"os"
	"text/tabwriter"
)

func PrintOrganzationFabricNodes(runtimeFabric *entities.OrganizationFabricEntity) {

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s\t%s\t%s", "NAME", "READY", "HEALTHY", "SCHEDULABLE", "KUBELET", "DOCKER", "ROLE", "CAPACITY")

	for _, v := range runtimeFabric.Nodes {

		nodeSummary := builders.NewRuntimeFabricNodeSummaryBuilder(&[]response.FabricNode{v}).Build()
		fmt.Fprintf(w, "\n %s\t%v\t%v\t%v\t%s\t%s\t%s\t%s",
			v.Name,
			v.Status.IsReady,
			v.Status.IsHealthy,
			v.Status.IsSchedulable,
			v.KubeletVersion,
			v.DockerVersion,
			v.Role,
			nodeSummary.Capacity,
		)
	}

	fmt.Fprintf(w, "\n")
}

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
