package cmd

import (
	"fmt"
	"os"

	"amrita_pyq/cmd/internal/configs"
	"amrita_pyq/cmd/internal/logo"
	"amrita_pyq/cmd/internal/menu"
	"amrita_pyq/cmd/internal/requestclient"
	"amrita_pyq/cmd/internal/webclient"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "ampyq",
	Short: "Amrita PYQ CLI",
	Long:  `A CLI application to access Amrita Repository for previous year question papers.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(configs.LogoStyle.Render(logo.LOGO_ASCII))
		cs := menu.CourseSelect{
			ReqClient: requestclient.RequestClient{
				WebClient: webclient.DefaultWebClient{},
			},
		}
		cs.ChooseCourse()
	},
}

func Execute() {
	if err := Cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
