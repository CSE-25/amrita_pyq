package version

import (
	"amrita_pyq/cmd/root"

	"github.com/spf13/cobra"
)

func init() {
	root.RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of ampyq",
	Long:  `Displays version of ampyq installed on the system.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Amrita Previous Year Questions v0.0.1-alpha")
	},
}
