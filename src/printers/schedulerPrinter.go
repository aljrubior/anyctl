package printers

import (
	"fmt"
	"github.com/aljrubior/anyctl/managers/entities"
	"os"
	"strconv"
	"text/tabwriter"
)

func NewSchedulerPrinter(entity *entities.SchedulerEntity) *SchedulerPrinter {
	return &SchedulerPrinter{
		entity: entity,
	}
}

type SchedulerPrinter struct {
	entity *entities.SchedulerEntity
}

func (this *SchedulerPrinter) Print() {
	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	if this.entity.Type == "CronScheduler" {
		fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s", "FLOW NAME", "TYPE", "ENABLED", "EXPRESSION", "TIME ZONE")
		fmt.Fprintf(w, fmt.Sprintf("\n %s\t%s\t%s\t%s\t%s \n",
			this.entity.FlowName,
			this.entity.Type,
			strconv.FormatBool(this.entity.Enabled),
			this.entity.Expression,
			this.entity.TimeZone))
		return
	}

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s", "FLOW NAME", "TYPE", "ENABLED", "DELAY", "FREQUENCY", "TIME UNIT")

	fmt.Fprintf(w, fmt.Sprintf("\n %s\t%s\t%s\t%s\t%s\t%s \n",
		this.entity.FlowName,
		this.entity.Type,
		strconv.FormatBool(this.entity.Enabled),
		strconv.Itoa(this.entity.StartDelay),
		strconv.Itoa(this.entity.Frequency),
		this.entity.TimeUnit))
}
