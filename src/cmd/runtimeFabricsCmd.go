package cmd

import (
	"github.com/spf13/cobra"
)

var runtimeFabricsCmd = &cobra.Command{
	Use:     "runtimefabrics",
	Aliases: []string{"runtimefabric", "rtf"},
	Short:   "Runtime Fabric resource",
	Run: func(cmd *cobra.Command, args []string) {

		println("Error: Unsupported option. Try with 'anyctl runtimemanager runtimefabrics --help'")
	},
}

func init() {
	runtimeManagerCmd.AddCommand(runtimeFabricsCmd)
}
