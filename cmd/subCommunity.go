package cmd

import (
	"fmt"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"os"
	"strings"
	"time"
)

type SubCommunity struct {
	name string
	path string
}

func subCommTable(url string) {
	for {
		action := func() {
			time.Sleep(2 * time.Second)
		}
		if err := spinner.New().Title("Fetching Sub Communities...").Action(action).Run(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		subCommunities, err := subComReq(url)

		if err != nil {
			fmt.Print(errorStyle.Render(fmt.Sprintf("Error: %v\n", err)))
			return
		}

		var selectedOption string
		var options []huh.Option[string]
		var subCommunityList []SubCommunity

		// Convert sub communities to huh options.
		for _, subCom := range subCommunities {
			subCommItem := SubCommunity(subCom)
			subCommunityList = append(subCommunityList, subCommItem)
			options = append(options, huh.NewOption(subCommItem.name, subCommItem.path))
		}
		// Add back option.
		options = append(options, huh.NewOption("Back to Main Menu", "back"))
		options = append(options, huh.NewOption("Quit", "quit"))
		selectionDisplay := "Selection(s):\n" + strings.Join(selectionHistory, " → ")

		// Create the select field.
		selectField := huh.NewSelect[string]().
			Title("Select Sub community to view").
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
					TitleFunc(func() string { return selectionDisplay }, &selectionHistory),
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
			huhMenuStart() // Go back to main menu.
		case "quit":
			QuitWithSpinner()
		default:
			// Find selected sub Community and process it
			for _, subComItem := range subCommunityList {
				if subComItem.path == selectedOption {
					url := BASE_URL + subComItem.path
					yearTable(url) // Function to display the question papers
					break
				}
			}
		}
	}
}
