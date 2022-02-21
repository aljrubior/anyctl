package printers

import (
	"fmt"
	"github.com/aljrubior/anyctl/managers/entities"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

func NewSchedulersPrinter(entities *[]entities.SchedulerEntity) *SchedulersPrinter {

	return &SchedulersPrinter{
		entities: entities,
	}
}

type SchedulersPrinter struct {
	entities *[]entities.SchedulerEntity
}

func (this *SchedulersPrinter) Print() {
	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 0, 0, 3, ' ', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s", "FLOW NAME", "TYPE", "ENABLED", "SCHEDULE")

	for _, v := range *this.entities {
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
