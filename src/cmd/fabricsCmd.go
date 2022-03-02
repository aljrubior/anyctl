package cmd

import (
	"github.com/spf13/cobra"
)

var fabricsCmd = &cobra.Command{
	Use:     "fabrics",
	Aliases: []string{"fabric", "fa"},
	Short:   "Fabrics resource",
	Run: func(cmd *cobra.Command, args []string) {

		println("Error: Unsupported option. Try with 'anyctl fabrics --help'")
	},
}

func init() {
	adminCmd.AddCommand(fabricsCmd)
}
