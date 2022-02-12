package builders

import (
	"fmt"
	"github.com/aljrubior/anyctl/clients/fabrics/response"
	"github.com/aljrubior/anyctl/model"
)

func NewRuntimeFabricNodeSummaryBuilder(runtimeFabricNodes *[]response.FabricNode) *RuntimeFabricNodeSummaryBuilder {
	return &RuntimeFabricNodeSummaryBuilder{
		runtimeFabricNodes,
	}
}

type RuntimeFabricNodeSummaryBuilder struct {
	runtimeFabricNodes *[]response.FabricNode
}

func (this *RuntimeFabricNodeSummaryBuilder) Build() *model.RuntimeFabricNodeSummary {
	healthyCount := 0
	readyCount := 0
	schedulableCount := 0
	capacityMemoryCount := 0
	capacityCpuCount := 0
	capacityPodCount := 0

	requestMemoryCount := 0
	requestCpuCount := 0
	requestPodCount := 0

	limitMemoryCount := 0
	limitCpuCount := 0
	limitPodCount := 0

	for _, v := range *this.runtimeFabricNodes {

		if v.Status.IsHealthy {
			healthyCount++
		}

		if v.Status.IsReady {
			readyCount++
		}

		if v.Status.IsSchedulable {
			schedulableCount++
		}

		capacityMemoryCount += v.Capacity.MemoryMi
		capacityCpuCount += v.Capacity.Cpu
		capacityPodCount += v.Capacity.Pods

		requestMemoryCount += v.AllocatedRequestCapacity.MemoryMi
		requestCpuCount += v.AllocatedRequestCapacity.Cpu
		requestPodCount += v.AllocatedRequestCapacity.Pods

		limitMemoryCount += v.AllocatedLimitCapacity.MemoryMi
		limitCpuCount += v.AllocatedLimitCapacity.Cpu
		limitPodCount += v.AllocatedLimitCapacity.Pods
	}

	total := len(*this.runtimeFabricNodes)

	return &model.RuntimeFabricNodeSummary{
		Healty:          fmt.Sprintf("%d/%d", healthyCount, total),
		Ready:           fmt.Sprintf("%d/%d", readyCount, total),
		Schedulable:     fmt.Sprintf("%d/%d", schedulableCount, total),
		Capacity:        fmt.Sprintf("%d Mem / %d Cpu / %d Pods", capacityMemoryCount, capacityCpuCount, capacityPodCount),
		RequestCapacity: fmt.Sprintf("%d Mem / %d Cpu / %d Pods", requestMemoryCount, requestCpuCount, requestPodCount),
		LimitCapacity:   fmt.Sprintf("%d Mem / %d Cpu / %d Pods", limitMemoryCount, limitCpuCount, limitPodCount),
	}
}
