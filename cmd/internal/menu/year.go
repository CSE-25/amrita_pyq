package menu

import (
	"fmt"
	"os"
	"strings"
	"time"

	"amrita_pyq/cmd/internal/configs"
	"amrita_pyq/cmd/internal/requestclient"
	"amrita_pyq/cmd/util"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
)

type (
	File struct {
		name string
		path string
	}

	YearTable struct {
		ReqClient requestclient.RequestClient
	}
)

func (yt *YearTable) ChooseQP(url string) {
	for {
		action := func() {
			time.Sleep(2 * time.Second)
		}
		if err := spinner.New().Title("Fetching ...").Action(action).Run(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		files, err := yt.ReqClient.YearReq(url)
		if err != nil {
			fmt.Print(configs.ErrorStyle.Render(fmt.Sprintf("Error: %v\n", err)))
			return
		}

		var selectedOption string
		var options []huh.Option[string]
		var fileList []File

		// Add back option.
		options = append(options, huh.NewOption("Back to Main Menu", "back"))
		options = append(options, huh.NewOption("Quit", "quit"))

		// Convert files to huh options.
		for _, file := range files {
			fileItem := File{file.Name, file.Path}
			fileList = append(fileList, fileItem)
			options = append(options, huh.NewOption(fileItem.name, fileItem.path))
		}

		selectionDisplay := "Selection(s):\n" + strings.Join(configs.SelectionHistory, " → ")

		// Create the select field.
		selectField := huh.NewSelect[string]().
			Title("Select Question Paper to view").
			Options(options...).
			Value(&selectedOption)

		// Limit number of options displayed when
		// there are more than 10 options to prevent overflow.
		if len(options) > 10 {
			selectField = selectField.WithHeight(20).(*huh.Select[string])
		}

		// Create the form.
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewNote().
					TitleFunc(func() string { return selectionDisplay }, &configs.SelectionHistory),
				selectField,
			),
		)

		// Run the form.
		err = form.Run()
		if err != nil {
			fmt.Printf("Error: %v", err)
			os.Exit(1)
		}

		// Handle selection.
		switch selectedOption {
		case "back":
			cs := CourseSelect{
				ReqClient: yt.ReqClient,
			}
			cs.ChooseCourse()
		case "quit":
			util.QuitWithSpinner() // Quit the program.
		default:
			// Find selected file and process it.
			for _, fileItem := range fileList {
				if fileItem.path == selectedOption {
					url := configs.BASE_URL + fileItem.path

					// Open the file in the browser.
					if err := yt.ReqClient.WebClient.OpenBrowser(url); err != nil {
						fmt.Println(configs.ErrorStyle.Render(fmt.Sprintf("Error: %v\n", err)))
					}
					break
				}
			}
		}
	}
}
