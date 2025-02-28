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
	Assessment struct {
		name string
		path string
	}

	SemChoose struct {
		ReqClient requestclient.RequestClient
	}
)

func (sc *SemChoose) ChooseSemester(url string) {
	action := func() {
		time.Sleep(2 * time.Second)
	}
	if err := spinner.New().Title("Fetching assessments").Action(action).Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	params_url := url

	assessments, err := sc.ReqClient.SemChooseReq(url)
	if err != nil {
		fmt.Println(configs.ErrorStyle.Render(fmt.Sprintf("Error: %v\n", err)))
		return
	}

	var selectedOption string
	var assessList []Assessment
	var options []huh.Option[string]
	
	// Convert assessments to huh options.
	for _, assessment := range assessments {
		assess := Assessment{assessment.Name, assessment.Path}
		assessList = append(assessList, assess)
		options = append(options, huh.NewOption(assess.name, assess.name))
	}

	// Add back and quit option.
	options = append(options, huh.NewOption("Back", "Back"))
	options = append(options, huh.NewOption("Quit", "Quit"))
	
	selectionDisplay := "Selection(s):\n" + strings.Join(configs.SelectionHistory, " → ")
	// Create the form.
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewNote().
				TitleFunc(func() string { return selectionDisplay }, &configs.SelectionHistory),
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

	if selectedOption == "Back" && len(configs.SelectionHistory) > 0 {
		configs.SelectionHistory = configs.SelectionHistory[:len(configs.SelectionHistory)-1] // Remove last selection
	} else {
		configs.SelectionHistory = append(configs.SelectionHistory, selectedOption) // Append new selection
	}

	// Handle selection.
	if selectedOption == "Back" {
		semTable := SemTable{
			ReqClient: sc.ReqClient,
		}
		semTable.ChooseQuestionSetFromSemester(stack.STACK.Pop()) // Proper call to previous menu
		return
	}

	// Auto-exit if "Quit" is selected
	if selectedOption == "Quit" {
		util.QuitWithSpinner()
	}

	// Find selected assessment and process it.
	for _, assess := range assessList {
		if assess.name == selectedOption {
			url := configs.BASE_URL + assess.path
			subComTable := SubCommunityTable{
				ReqClient: sc.ReqClient,
			}
			subComTable.ChooseSubCommunity(url)
			break
		}
	}

	stack.STACK.Push(params_url)
}
