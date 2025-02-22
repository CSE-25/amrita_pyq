package version

import (
	"amrita_pyq/cmd"
	"github.com/spf13/cobra"
)

func init() {
	cmd.Cmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of ampyq",
	Long:  `Displays version of ampyq installed on the system.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Amrita Previous Year Questions v1.0.1")
	},
}
