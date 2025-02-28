package menu

import (
	"fmt"
	"os"
	"strings"
	"time"

	"amrita_pyq/cmd/internal/configs"
	"amrita_pyq/cmd/internal/requestclient"
	"amrita_pyq/cmd/util"
	"amrita_pyq/cmd/util/stack"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
)

type (
	Semester struct {
		name string
		path string
	}

	SemTable struct {
		ReqClient requestclient.RequestClient
	}
)

func (sc *SemTable) ChooseQuestionSetFromSemester(url string) {
	action := func() {
		time.Sleep(2 * time.Second)
	}
	if err := spinner.New().Title("Fetching Semesters").Action(action).Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	semesters, err := sc.ReqClient.SemTableReq(url)
	if err != nil {
		fmt.Print(configs.ErrorStyle.Render(fmt.Sprintf("Error: %v\n", err)))
		return
	}

	var selectedOption string
	var sems []Semester
	var options []huh.Option[string]

	// Add back and quit option.
	options = append(options, huh.NewOption("Back", "Back"))
	options = append(options, huh.NewOption("Quit", "Quit"))

	// Convert semesters to huh options.
	for _, sem := range semesters {
		semester := Semester{sem.Name, sem.Path}
		sems = append(sems, semester)
		options = append(options, huh.NewOption(semester.name, semester.name))
	}

	selectionDisplay := "Selection(s):\n" + strings.Join(configs.SelectionHistory, " → ")
	// Create the form.
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewNote().
				TitleFunc(func() string { return selectionDisplay }, &configs.SelectionHistory),
			huh.NewSelect[string]().
				Title("Semesters").
				Options(options...).
				Value(&selectedOption),
		),
	)

	stack.STACK.Push(url) // Save current URL to stack.

	err = form.Run()
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
	if selectedOption == "Back" && len(configs.SelectionHistory) > 0 {
		configs.SelectionHistory = configs.SelectionHistory[:len(configs.SelectionHistory)-1] // Remove last selection
	} else {
		configs.SelectionHistory = append(configs.SelectionHistory, selectedOption) // Append new selection
	}

	// Handle selection.
	if selectedOption == "Back" {
		cs := CourseSelect{
			ReqClient: sc.ReqClient,
		}
		cs.ChooseCourse()
		return
	}

	// Auto-exit if "Quit" is selected
	if selectedOption == "Quit" {
		util.QuitWithSpinner()
	}

	// Find selected semester and process it.
	for _, sem := range sems {
		if sem.name == selectedOption {
			url := configs.BASE_URL + sem.path
			semChoose := SemChoose{
				ReqClient: sc.ReqClient,
			}
			semChoose.ChooseSemester(url)
			break
		}
	}
}
