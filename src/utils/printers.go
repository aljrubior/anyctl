package utils

import (
	"fmt"
	"github.com/aljrubior/anyctl/builders"
	"github.com/aljrubior/anyctl/clients/fabrics/response"
	"github.com/aljrubior/anyctl/managers/entities"
	"github.com/aljrubior/anyctl/managers/wrappers"
	"github.com/aljrubior/anyctl/manifests"
	"gopkg.in/yaml.v2"
	"os"
	"text/tabwriter"
)

func PrintDeployments(deployments *[]entities.DeploymentItemEntity, targets *[]entities.TargetEntity) {

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	targetsMap := TargetEntities2Map(*targets)

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s", "NAME", "DSTATUS", "ASTATUS", "TARGET")

	for _, v := range *deployments {
		fmt.Fprintf(w, "\n %s\t%s\t%s\t%s",
			v.Name,
			v.Status,
			v.Application.Status,
			targetsMap[v.Target.TargetId].GetName())
	}

	fmt.Fprintf(w, "\n")
}

func PrintDeployment(deployment *entities.DeploymentEntity, targets *[]entities.TargetEntity) {

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	targetsMap := TargetEntities2Map(*targets)

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s", "NAME", "REPLICAS", "STATUS", "TARGET", "RUNTIME", "ASSET")

	fmt.Fprintln(w, fmt.Sprintf("\n %s\t%s\t%s\t%s\t%s\t%s",
		deployment.Name,
		BuildReplicasSummary(deployment.Replicas),
		deployment.Status,
		targetsMap[deployment.Target.TargetId].GetName(),
		deployment.Target.DeploymentSettings.RuntimeVersion,
		fmt.Sprintf("%s:%s", deployment.Application.Ref.ArtifactId, deployment.Application.Ref.Version)))
}

func PrintAssets(assets *[]entities.AssetEntity) {
	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s", "NAME", "VERSION", "STATUS", "RUNTIME", "UPDATED AT", "CREATED BY")

	for _, v := range *assets {
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

func PrintTargets(targets *[]entities.TargetEntity) {

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s", "ID", "NAME", "TYPE")

	for _, v := range *targets {
		fmt.Fprintf(w, "\n %s\t%s\t%s",
			v.GetId(),
			v.GetName(),
			v.GetType())
	}

	fmt.Fprintf(w, "\n")
}

func PrintTargetSupportedVersions(targetWrapper *wrappers.TargetEntityWrapper) {

	target, ok := targetWrapper.GetRuntimeFabricTargetEntity()

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

func PrintStandaloneDetails(targetWrapper *wrappers.TargetEntityWrapper) {

	target, ok := targetWrapper.GetStandaloneTargetEntity()

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

func PrintStandaloneAddresses(targetWrapper *wrappers.TargetEntityWrapper) {

	target, ok := targetWrapper.GetStandaloneTargetEntity()

	if !ok {
		println("Addresses is not available on this target type")
		return
	}

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s", "INTERFACE", "IP")

	for _, v := range target.Details.Addresses {
		fmt.Fprintf(w, "\n %s\t%s",
			v.NetworkInterface,
			v.Ip)
	}

	fmt.Fprintf(w, "\n")
}

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

func PrintSharedSpaces(sharedSpaces *[]entities.SharedSpaceEntity) {

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s", "NAME", "REGION", "FLAVOR", "STATUS")

	for _, v := range *sharedSpaces {
		fmt.Fprintf(w, "\n %s\t%s\t%s\t%s",
			v.Name,
			v.Region,
			v.Flavor,
			v.Status,
		)
	}

	fmt.Fprintf(w, "\n")
}

func PrintSharedSpace(sharedSpace *entities.SharedSpaceEntity, privateSpace *entities.PrivateSpaceEntity) {

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s\t%s", "NAME", "REGION", "FLAVOR", "STATUS", "ADVERTISED", "REQUIRES PERMISSION", "PRIVATE SPACE")

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%v\t%v\t%s",
		sharedSpace.Name,
		sharedSpace.Region,
		sharedSpace.Flavor,
		sharedSpace.Status,
		sharedSpace.IsAdvertised,
		sharedSpace.RequiresPermission,
		privateSpace.Name,
	)

	fmt.Fprintf(w, "\n")
}
