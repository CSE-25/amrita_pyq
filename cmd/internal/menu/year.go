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

		// Convert files to huh options.
		for _, file := range files {
			fileItem := File{file.Name, file.Path}
			fileList = append(fileList, fileItem)
			options = append(options, huh.NewOption(fileItem.name, fileItem.path))
		}

		// Add back and quit option.
		options = append(options, huh.NewOption("Back to Main Menu", "back"))
		options = append(options, huh.NewOption("Quit", "quit"))

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
			for _, fileItem := range fileList {
				if fileItem.path == selectedOption {
					url := configs.BASE_URL + fileItem.path

					var actionChoice string
					choiceOptions := []huh.Option[string]{
						huh.NewOption("Open in Browser", "open"),
						huh.NewOption("Download", "download"),
					}

					choiceField := huh.NewSelect[string]().
						Title("Choose an action").
						Options(choiceOptions...).
						Value(&actionChoice)

					choiceForm := huh.NewForm(
						huh.NewGroup(choiceField),
					)
					choiceForm.Run()

					switch actionChoice {
					case "open":
						if err := yt.ReqClient.WebClient.OpenBrowser(url); err != nil {
							fmt.Println(configs.ErrorStyle.Render(fmt.Sprintf("Error: %v\n", err)))
						}
					case "download":
						action := func() {
							time.Sleep(2 * time.Second)
						}
						if err := spinner.New().Type(spinner.Dots).Title("Downloading ...").Action(action).Run(); err != nil {
							fmt.Println(err)
							os.Exit(1)
						}
						if err := yt.ReqClient.WebClient.DownloadFile(url, fileItem.name); err != nil {
							fmt.Println(configs.ErrorStyle.Render(fmt.Sprintf("Download Error: %v\n", err)))
						}
						fmt.Println(configs.FetchStatusStyle.Render(fmt.Sprintf("Download complete: amritapyq/pyq_downloads/%s", fileItem.name)))
					}
					break
				}
			}
		}
	}
}
