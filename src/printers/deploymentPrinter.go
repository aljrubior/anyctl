package printers

import (
	"fmt"
	"github.com/aljrubior/anyctl/clients/deployments/response"
	"github.com/aljrubior/anyctl/managers/entities"
	"os"
	"strconv"
	"text/tabwriter"
)

func NewDeploymentPrinter(entity *entities.DeploymentEntity, targets *[]entities.TargetEntity) *DeploymentPrinter {

	return &DeploymentPrinter{
		entity:  entity,
		targets: targets,
	}
}

type DeploymentPrinter struct {
	entity  *entities.DeploymentEntity
	targets *[]entities.TargetEntity
}

func (this DeploymentPrinter) Print() {

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	targetsMap := this.transformTargetEntitiesToMap(*this.targets)

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s", "NAME", "REPLICAS", "STATUS", "TARGET", "RUNTIME", "ASSET")

	fmt.Fprintln(w, fmt.Sprintf("\n %s\t%s\t%s\t%s\t%s\t%s",
		this.entity.Name,
		this.buildReplicasSummary(this.entity.Replicas),
		this.entity.Status,
		targetsMap[this.entity.Target.TargetId].GetName(),
		this.entity.Target.DeploymentSettings.RuntimeVersion,
		fmt.Sprintf("%s:%s", this.entity.Application.Ref.ArtifactId, this.entity.Application.Ref.Version)))
}

func (this *DeploymentPrinter) transformTargetEntitiesToMap(fromTargets []entities.TargetEntity) map[string]entities.TargetEntity {

	result := make(map[string]entities.TargetEntity)

	for _, v := range fromTargets {
		result[v.GetId()] = v
	}

	return result
}

func (this *DeploymentPrinter) buildReplicasSummary(fromReplicas []response.Replica) string {
	total := len(fromReplicas)
	started := 0

	for _, v := range fromReplicas {
		if v.State == "STARTED" {
			started++
		}
	}

	return fmt.Sprintf("%s/%s", strconv.Itoa(started), strconv.Itoa(total))
}
