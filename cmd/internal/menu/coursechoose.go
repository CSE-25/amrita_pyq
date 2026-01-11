package menu

import (
	"fmt"
	"os"
	"time"

	"amrita_pyq/cmd/internal/configs"
	"amrita_pyq/cmd/internal/requestclient"
	"amrita_pyq/cmd/util"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
)

type (
	Subject struct {
		name string
		path string
	}

	CourseSelect struct {
		ReqClient requestclient.RequestClient
	}
)

func (cs *CourseSelect) ChooseCourse() {
	action := func() {
		time.Sleep(2 * time.Second)
	}
	if err := spinner.New().Title("Fetching Courses").Action(action).Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	resources, err := cs.ReqClient.GetCoursesReq(configs.COURSE_LIST_URL)
	if err != nil {
		fmt.Println(configs.ErrorStyle.Render(fmt.Sprintf("Error: %v\n", err)))
		os.Exit(1)
	}

	var selectedOption string
	var subjects []Subject
	var options []huh.Option[string]

	for _, res := range resources {
		subject := Subject{res.Name, res.Path}
		subjects = append(subjects, subject)
		options = append(options, huh.NewOption(subject.name, subject.name))
	}

	// --- Added Go Back and Quit ---
	options = append(options, huh.NewOption("Go Back", "Back"))
	options = append(options, huh.NewOption("Quit", "Quit"))

	// First menu does NOT display history yet
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Available Courses").
				Options(options...).
				Value(&selectedOption),
		),
	)

	err = form.Run()
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	// Auto-exit if "Quit" or "Back" is selected
	if selectedOption == "Quit" || selectedOption == "Back" {
		util.QuitWithSpinner()
	}

	// Store only the selected course
	configs.SelectionHistory = []string{selectedOption} // Reset history to show only last selected

	// Move to the next menu (Second Menu)
	for _, subject := range subjects {
		if subject.name == selectedOption {
			url := configs.BASE_URL + subject.path
			semTable := SemTable{
				ReqClient: cs.ReqClient,
			}
			semTable.ChooseQuestionSetFromSemester(url)
			break
		}
	}
}
