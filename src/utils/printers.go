package utils

import (
	"encoding/json"
	"fmt"
	"github.com/aljrubior/anyctl/builders"
	"github.com/aljrubior/anyctl/clients/fabrics/response"
	"github.com/aljrubior/anyctl/comparators"
	"github.com/aljrubior/anyctl/managers/entities"
	"github.com/aljrubior/anyctl/managers/wrappers"
	"github.com/aljrubior/anyctl/manifests"
	"gopkg.in/yaml.v2"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
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
		fmt.Sprintf("%s:%s", deployment.Application.Asset.ArtifactId, deployment.Application.Asset.Version)))
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

func PrintPrivateSpace(privateSpace *entities.PrivateSpaceEntity) {

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s", "NAME", "REGION", "STATUS", "VERSION", "FLAVOR", "ENVIRONMENT TYPE")

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%v\t%v",
		privateSpace.Name,
		privateSpace.Region,
		privateSpace.Status,
		privateSpace.Version,
		privateSpace.Flavor,
		privateSpace.Environments.Type,
	)

	fmt.Fprintf(w, "\n")
}

func PrintPrivateSpaces(privateSpaces *[]entities.PrivateSpaceEntity) {

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s", "NAME", "REGION", "STATUS", "VERSION", "FLAVOR", "ENVIRONMENT TYPE")

	for _, v := range *privateSpaces {
		fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%v\t%v",
			v.Name,
			v.Region,
			v.Status,
			v.Version,
			v.Flavor,
			v.Environments.Type)
	}

	fmt.Fprintf(w, "\n")
}

func PrintPrivateSpaceFabrics(fabrics *[]entities.FabricEntity) {

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s\t%s", "NAME", "REGION", "VERSION", "STATUS", "LEVEL", "INFRA VERSION", "INFRA ID")

	for _, v := range *fabrics {
		fmt.Fprintf(w, "\n  %s\t%s\t%s\t%s\t%s\t%s\t%s",
			v.Name,
			v.Region,
			v.Version,
			v.Status,
			v.ClusterConfigurationLevel,
			v.InfraVersion,
			v.InfraDeploymentId)
	}

	fmt.Fprintf(w, "\n")
}

func PrintFabrics(fabrics *[]entities.FabricEntity, organizations map[string]*entities.OrganizationEntity) {

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s", "NAME", "ORG NAME", "ORG TYPE", "VERSION", "REGION", "STATUS")

	for _, v := range *fabrics {

		fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s",
			v.Name,
			organizations[v.OrganizationId].Name,
			organizations[v.OrganizationId].Subscription.Type,
			v.Version,
			v.Region,
			v.Status)
	}

	fmt.Fprintf(w, "\n")
}

func substr(value string, index, length int) string {

	if value == "" || len(value) < length {
		return value
	}

	return value[index:length]
}

func PrintOrganizationPrivateSpace(privateSpace *entities.OrganizationPrivateSpaceEntity) {

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s", "NAME", "REGION", "STATUS", "DNS")

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s",
		privateSpace.Name,
		privateSpace.Region,
		privateSpace.Status,
		privateSpace.Network.DnsTarget,
	)

	fmt.Fprintf(w, "\n")
}

func PrintOrganizationPrivateSpaceFirewallRules(privateSpace *entities.OrganizationPrivateSpaceEntity) {

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s", "CIDR BLOCK", "PROTOCOL", "FROM PORT", "TO PORT", "TYPE")

	for _, v := range privateSpace.FirewallRules {
		fmt.Fprintf(w, "\n %s\t%s\t%d\t%d\t%s",
			v.CidrBlock,
			v.Protocol,
			v.FromPort,
			v.ToPort,
			v.Type,
		)
	}

	fmt.Fprintf(w, "\n")
}

func PrintPrivateSpaceManagedFirewallRules(privateSpace *entities.PrivateSpaceEntity) {

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s", "CIDR BLOCK", "PROTOCOL", "FROM PORT", "TO PORT", "TYPE")

	for _, v := range privateSpace.ManagedFirewallRules {
		fmt.Fprintf(w, "\n %s\t%s\t%d\t%d\t%s",
			v.CidrBlock,
			v.Protocol,
			v.FromPort,
			v.ToPort,
			v.Type,
		)
	}

	fmt.Fprintf(w, "\n")
}

func PrintOrganizationPrivateSpaceNetwork(privateSpace *entities.OrganizationPrivateSpaceEntity) {

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s", "REGION", "CIDR BLOCK", "INBOUND IPS", "OUTBOUND IPS", "DNS")

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s",
		privateSpace.Network.Region,
		privateSpace.Network.CidrBlock,
		strings.Join(privateSpace.Network.InboundStaticIps, ", "),
		strings.Join(privateSpace.Network.OutboundStaticIps, ", "),
		privateSpace.Network.DnsTarget,
	)

	fmt.Fprintf(w, "\n")
}

func PrintOrganizationPrivateSpaceManifest(manifest *manifests.OrganizationPrivateSpaceManifest) {

	data, err := yaml.Marshal(*manifest)

	if err != nil {
		println(err)
		return
	}

	println(string(data))
}

func PrintPrivateSpaceManifest(manifest *manifests.PrivateSpaceManifest) {

	data, err := yaml.Marshal(*manifest)

	if err != nil {
		println(err)
		return
	}

	println(string(data))
}

func PrintFabricManifest(manifest *manifests.FabricManifest) {

	data, err := yaml.Marshal(*manifest)

	if err != nil {
		println(err)
		return
	}

	println(string(data))
}

func PrintFabric(fabric *entities.FabricEntity) {

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s", "NAME", "REGION", "VERSION", "STATUS", "LEVEL", "INFRA ID")

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s",
		fabric.Name,
		fabric.Region,
		fabric.Version,
		fabric.Status,
		fabric.ClusterConfigurationLevel,
		fabric.InfraDeploymentId)

	fmt.Fprintf(w, "\n")

}

func PrintFabricVersionInformation(fabric *entities.FabricEntity) {

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s", "VERSION", "DESIRED", "UPGRADE AVAILABLE", "KUBERNETES", "INFRA", "DESIRED INFRA")

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s",
		fabric.Version,
		fabric.DesiredVersion,
		fabric.AvailableUpgradeVersion,
		fabric.KubernetesVersion,
		fabric.InfraVersion,
		fabric.DesiredInfraVersion)

	fmt.Fprintf(w, "\n")

}

func PrintSharedSpaceManifest(manifest *manifests.SharedSpaceManifest) {

	data, err := yaml.Marshal(*manifest)

	if err != nil {
		println(err)
		return
	}

	println(string(data))
}

func PrintSchedulers(schedulers *[]entities.SchedulerEntity) {

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s", "FLOW NAME", "TYPE", "ENABLED", "SCHEDULE")

	for _, v := range *schedulers {
		if v.Type == "CronScheduler" {
			fmt.Fprintf(w, fmt.Sprintf("\n %s\t%s\t%s\t%s",
				v.FlowName,
				v.Type,
				strconv.FormatBool(v.Enabled),
				v.Expression))
			continue
		}

		fmt.Fprintf(w, fmt.Sprintf("\n %s\t%s\t%s\t%s",
			v.FlowName,
			v.Type,
			strconv.FormatBool(v.Enabled),
			fmt.Sprintf("Every %d %s", v.Frequency, strings.ToLower(v.TimeUnit))))

	}

	fmt.Fprintf(w, "\n")
}

func PrintScheduler(scheduler *entities.SchedulerEntity) {

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	if scheduler.Type == "CronScheduler" {
		fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s", "FLOW NAME", "TYPE", "ENABLED", "EXPRESSION", "TIME ZONE")
		fmt.Fprintf(w, fmt.Sprintf("\n %s\t%s\t%s\t%s\t%s \n",
			scheduler.FlowName,
			scheduler.Type,
			strconv.FormatBool(scheduler.Enabled),
			scheduler.Expression,
			scheduler.TimeZone))
		return
	}

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s", "FLOW NAME", "TYPE", "ENABLED", "DELAY", "FREQUENCY", "TIME UNIT")

	fmt.Fprintf(w, fmt.Sprintf("\n %s\t%s\t%s\t%s\t%s\t%s \n",
		scheduler.FlowName,
		scheduler.Type,
		strconv.FormatBool(scheduler.Enabled),
		strconv.Itoa(scheduler.StartDelay),
		strconv.Itoa(scheduler.Frequency),
		scheduler.TimeUnit))
}

func PrintOrgUsage(orgs []*entities.OrganizationEntity) {

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s",
		"ORG ID",
		"ORG NAME",
		"PROD VCORES",
		"SAND VCORES",
		"DESIGN VCORES",
		"STATIC IPS",
		"VPCS",
		"VPNS",
		"LB")

	for _, v := range orgs {

		if shouldPrintOrgUsage(v) {
			fmt.Fprintf(w, fmt.Sprintf("\n %s\t%s\t%v\t%v\t%v\t%v\t%v\t%v\t%v",
				v.Id,
				v.Name,
				formatFloat(v.Entitlements.VCoresProduction.Reassigned),
				formatFloat(v.Entitlements.VCoresSandbox.Reassigned),
				formatFloat(v.Entitlements.VCoresDesign.Reassigned),
				formatFloat(v.Entitlements.StaticIps.Reassigned),
				formatFloat(v.Entitlements.Vpcs.Reassigned),
				formatFloat(v.Entitlements.Vpns.Reassigned),
				formatFloat(v.Entitlements.LoadBalancer.Reassigned)))
		}
	}

	fmt.Fprintf(w, "\n")
}

func shouldPrintOrgUsage(org *entities.OrganizationEntity) bool {

	if org.OrganizationResponse.Entitlements.VCoresProduction.Reassigned > 0.0 ||
		org.OrganizationResponse.Entitlements.VCoresSandbox.Reassigned > 0.0 ||
		org.OrganizationResponse.Entitlements.VCoresDesign.Reassigned > 0.0 ||
		org.OrganizationResponse.Entitlements.StaticIps.Reassigned > 0.0 ||
		org.OrganizationResponse.Entitlements.Vpcs.Reassigned > 0.0 ||
		org.OrganizationResponse.Entitlements.Vpns.Reassigned > 0.0 ||
		org.OrganizationResponse.Entitlements.LoadBalancer.Reassigned > 0.0 {
		return true
	}

	return false
}

func formatFloat(value float64) string {
	if value > 0 {
		return fmt.Sprintf("%v", value)
	}

	return "-"
}

func PrintOrgQuotas(org *entities.OrganizationEntity) {

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s", "NAME", "ASSIGNED", "USAGE")

	fmt.Fprintf(w, fmt.Sprintf("\n %s\t%s\t%s", "Production vCores", formatFloat(org.Entitlements.VCoresProduction.Assigned), formatFloat(org.Entitlements.VCoresProduction.Reassigned)))
	fmt.Fprintf(w, fmt.Sprintf("\n %s\t%s\t%s", "Sandbox vCores", formatFloat(org.Entitlements.VCoresSandbox.Assigned), formatFloat(org.Entitlements.VCoresSandbox.Reassigned)))
	fmt.Fprintf(w, fmt.Sprintf("\n %s\t%s\t%s", "Design vCores", formatFloat(org.Entitlements.VCoresDesign.Assigned), formatFloat(org.Entitlements.VCoresDesign.Reassigned)))
	fmt.Fprintf(w, fmt.Sprintf("\n %s\t%s\t%s", "Static IPs", formatFloat(org.Entitlements.StaticIps.Assigned), formatFloat(org.Entitlements.StaticIps.Reassigned)))
	fmt.Fprintf(w, fmt.Sprintf("\n %s\t%s\t%s", "VPCs", formatFloat(org.Entitlements.Vpcs.Assigned), formatFloat(org.Entitlements.Vpcs.Reassigned)))
	fmt.Fprintf(w, fmt.Sprintf("\n %s\t%s\t%s", "VPNs", formatFloat(org.Entitlements.Vpns.Assigned), formatFloat(org.Entitlements.Vpns.Reassigned)))
	fmt.Fprintf(w, fmt.Sprintf("\n %s\t%s\t%s", "Load Balancers", formatFloat(org.Entitlements.LoadBalancer.Assigned), formatFloat(org.Entitlements.LoadBalancer.Reassigned)))

	fmt.Fprintf(w, "\n")
}

func PrintStandaloneTargetManifest(manifest *manifests.StandaloneTargetManifest) {

	data, err := yaml.Marshal(*manifest)

	if err != nil {
		println(err)
		return
	}

	println(string(data))
}

func PrintRuntimeFabricTargetManifest(manifest *manifests.RuntimeFabricTargetManifest) {

	data, err := yaml.Marshal(*manifest)

	if err != nil {
		println(err)
		return
	}

	println(string(data))
}

func PrintDeploymentManifest(manifest *manifests.DeploymentManifest) {

	data, err := yaml.Marshal(*manifest)

	if err != nil {
		println(err)
		return
	}

	println(string(data))
}

func PrintDeploymentSpecs(specs *[]entities.DeploymentSpecEntity) {
	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s", "DATE", "CHANGES")

	for i, v := range *specs {
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

func PrintDiffDeploymentSpecs(currentVersion, withVersion entities.DeploymentSpecEntity) error {
	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	currentVersionAsBytes, err := json.Marshal(currentVersion.DeploymentSpecResponse)

	if err != nil {
		return err
	}

	withVersionAsBytes, err := json.Marshal(withVersion.DeploymentSpecResponse)

	if err != nil {
		return err
	}

	var currentVersionAsMap map[interface{}]interface{}
	var withVersionAsMap map[interface{}]interface{}

	yaml.Unmarshal(currentVersionAsBytes, &currentVersionAsMap)
	yaml.Unmarshal(withVersionAsBytes, &withVersionAsMap)

	delete(currentVersionAsMap, "createdAt")
	delete(withVersionAsMap, "createdAt")

	differences := comparators.NewDeploymentComparator(currentVersionAsMap, withVersionAsMap).Compare()

	fmt.Fprintf(w, "Deployment differences are indicated with the following symbols:\n")
	fmt.Fprintf(w, "\t+ Add\n")
	fmt.Fprintf(w, "\t~ Change\n")
	fmt.Fprintf(w, "\t- Remove\n")
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "The following differences has been found between the current version '%s' and version '%s':\n", currentVersion.Version[:6], withVersion.Version[:6])
	fmt.Fprintf(w, "\n")

	var addCount, deleteCount, changeCount int

	for _, v := range differences {
		value := ""
		keyName := fmt.Sprintf("%v:", v.KeyName)

		switch v.Operator {
		case "~":
			if v.LeftValue == nil && v.RightValue == nil {
				value = ""
			} else {
				value = fmt.Sprintf("%s --> %v", v.LeftValue, v.RightValue)
				changeCount++
			}
		case "-":
			value = fmt.Sprintf("%v", v.RightValue)
			deleteCount++
		case "+":
			value = fmt.Sprintf("%v", v.RightValue)
			addCount++
		}

		println(valueOfDepth(v.Depth), v.Operator, keyName, value)
	}

	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "Result: %v to add, %v to change, %v to delete.", addCount, changeCount, deleteCount)
	fmt.Fprintf(w, "\n")

	return nil
}

func valueOfDepth(depth int) string {
	spaces := ""

	for i := 0; i <= depth; i++ {
		spaces += "  "
	}

	return spaces

}

func valueOf(value interface{}) string {
	if value == nil {
		return ""
	}

	return fmt.Sprintf("%v", value)
}
