package cmd

import (
	"fmt"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"os"
	"strings"
	"time"
)

type Assessment struct {
	name string
	path string
}

func semChoose(url string) {
	action := func() {
		time.Sleep(2 * time.Second)
	}
	if err := spinner.New().Title("Fetching assessments").Action(action).Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	params_url := url

	assessments, err := semChooseReq(url)
	if err != nil {
		fmt.Println(errorStyle.Render(fmt.Sprintf("Error: %v\n", err)))
		return
	}

	var selectedOption string
	var assessList []Assessment
	var options []huh.Option[string]

	// Convert assessments to huh options.
	for _, assessment := range assessments {
		assess := Assessment(assessment)
		assessList = append(assessList, assess)
		options = append(options, huh.NewOption(assess.name, assess.name))
	}
	// Add back and quit option.
	options = append(options, huh.NewOption("Back", "Back"))
	options = append(options, huh.NewOption("Quit", "Quit"))
	selectionDisplay := "Selection(s):\n" + strings.Join(selectionHistory, " → ")
	// Create the form.
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewNote().
				TitleFunc(func() string { return selectionDisplay }, &selectionHistory),
			huh.NewSelect[string]().
				Title("Assessments").
				Options(options...).
				Value(&selectedOption),
		),
	)

	err = form.Run()
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	if selectedOption == "Back" && len(selectionHistory) > 0 {
		selectionHistory = selectionHistory[:len(selectionHistory)-1] // Remove last selection
	} else {
		selectionHistory = append(selectionHistory, selectedOption) // Append new selection
	}

	// Handle selection.
	if selectedOption == "Back" {
		semTable(stack.Pop())
		return
	}

	// Auto-exit if "Quit" is selected
	if selectedOption == "Quit" {
		QuitWithSpinner()
	}

	// Find selected assessment and process it.
	for _, assess := range assessList {
		if assess.name == selectedOption {
			url := BASE_URL + assess.path
			year(url)
			break
		}
	}

	stack.Push(params_url)
}
