package year

import (
	"fmt"
	"os"
	"strings"
	"time"

	"amrita_pyq/cmd/helpers"
	"amrita_pyq/cmd/interfaces"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
)

// Interface to access functions from root package
var inter interfaces.Interface

func Init(n interfaces.Interface) {
	inter = n
}

type File struct {
	name string
	path string
}

func yearTable(url string) {
	for {
		action := func() {
			time.Sleep(2 * time.Second)
		}
		if err := spinner.New().Title("Fetching ...").Action(action).Run(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		files, err := inter.UseYearReq(url)
		if err != nil {
			fmt.Print(inter.UseErrorStyle().Render(fmt.Sprintf("Error: %v\n", err)))
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
		// Add back option.
		options = append(options, huh.NewOption("Back to Main Menu", "back"))
		options = append(options, huh.NewOption("Quit", "quit"))
		selectionDisplay := "Selection(s):\n" + strings.Join(helpers.SelectionHistory, " → ")

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
					TitleFunc(func() string { return selectionDisplay }, &helpers.SelectionHistory),
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
			inter.UseHuhMenuStart() // Go back to main menu.
		case "quit":
			inter.UseQuitWithSpinner() // Quit the program.
		default:
			// Find selected file and process it
			for _, fileItem := range fileList {
				if fileItem.path == selectedOption {
					url := inter.UseBASE_URL() + fileItem.path
					inter.UseOpenBrowser(url) // Function to open the browser with the selected URL.
					break
				}
			}
		}
	}
}

func Year(url string) {
	yearTable(url) // Call the yearTable function to display the menu.
}
