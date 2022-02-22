package utils

import (
	"fmt"
	"github.com/aljrubior/anyctl/builders"
	"github.com/aljrubior/anyctl/clients/fabrics/response"
	"github.com/aljrubior/anyctl/managers/entities"
	"github.com/aljrubior/anyctl/manifests"
	"gopkg.in/yaml.v2"
	"os"
	"text/tabwriter"
)

func PrintOrganizationFabrics(fabrics *[]entities.OrganizationFabricEntity) {

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s\t%s", "NAME", "REGION", "FABRIC VERSION", "STATUS", "AVAILABLE UPGRADE", "LEVEL", "DISTRIBUTION")

	for _, v := range *fabrics {
		fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s\t%s",
			v.Name,
			v.Region,
			v.Version,
			v.Status,
			v.AvailableUpgradeVersion,
			v.ClusterConfigurationLevel,
			getRuntimeFabricDistribution(&v),
		)
	}

	fmt.Fprintf(w, "\n")
}

func PrintOrganizationFabric(runtimeFabric *entities.OrganizationFabricEntity) {

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	nodeSummary := builders.NewRuntimeFabricNodeSummaryBuilder(&runtimeFabric.Nodes).Build()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s", "NAME", "READY", "HEALTHY", "SCHEDULABLE", "CAPACITY", "REGION", "VERSION", "STATUS", "DISTRIBUTION")

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s",
		runtimeFabric.Name,
		nodeSummary.Ready,
		nodeSummary.Healty,
		nodeSummary.Schedulable,
		nodeSummary.Capacity,
		runtimeFabric.Region,
		runtimeFabric.Version,
		runtimeFabric.Status,
		getRuntimeFabricDistribution(runtimeFabric),
	)

	fmt.Fprintf(w, "\n")
}

func PrintRuntimeFabricManifest(manifest *manifests.OrganizationFabricManifest) {

	data, err := yaml.Marshal(*manifest)

	if err != nil {
		println(err)
		return
	}

	println(string(data))
}

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

func getRuntimeFabricDistribution(runtimeFabric *entities.OrganizationFabricEntity) string {

	switch runtimeFabric.Vendor {
	case "aks":
		return "AKS"
	case "eks":
		return "EKS"
	case "gke":
		return "GKE"
	case "gravitational":
		return "APPLIANCE"
	case "rtfc":
		return "RTFC"
	default:
		return "Unknown"
	}
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
