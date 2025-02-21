package root

import (
	"fmt"
	"os"
	"time"

	"amrita_pyq/cmd/configs"
	"amrita_pyq/cmd/helpers"
	"amrita_pyq/cmd/logo"
	"amrita_pyq/cmd/model"
	"amrita_pyq/cmd/requestClient"
	"amrita_pyq/cmd/semChoose"
	"amrita_pyq/cmd/semTable"
	"amrita_pyq/cmd/year"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

// Used to implement the interface
type UseInterface struct{}

// Implementing from configs package
func (u *UseInterface) UseBASE_URL() string {
	return configs.BASE_URL
}
func (u *UseInterface) UseCOURSE_LIST_URL() string {
	return configs.COURSE_LIST_URL
}

// Implementing from helpers package
func (u *UseInterface) UseLogoStyle() lipgloss.Style {
	return helpers.LogoStyle
}
func (u *UseInterface) UseErrorStyle() lipgloss.Style {
	return helpers.ErrorStyle
}
func (u *UseInterface) UseFetchStatusStyle() lipgloss.Style {
	return helpers.FetchStatusStyle
}
func (u *UseInterface) UseFetchHTML(url string) (string, error) {
	return helpers.FetchHTML(url)
}
func (u *UseInterface) UseOpenBrowser(url string) error {
	return helpers.OpenBrowser(url)
}

// Implementing from logo package
func (u *UseInterface) UseLOGO_ASCII() string {
	return logo.LOGO_ASCII
}

// Implementing from requestClient package
func (u *UseInterface) UseGetCoursesReq(url string) ([]model.Resource, error) {
	return requestClient.GetCoursesReq(url)
}
func (u *UseInterface) UseSemChooseReq(url string) ([]model.Resource, error) {
	return requestClient.SemChooseReq(url)
}
func (u *UseInterface) UseSemTableReq(url string) ([]model.Resource, error) {
	return requestClient.SemTableReq(url)
}
func (u *UseInterface) UseYearReq(url string) ([]model.Resource, error) {
	return requestClient.YearReq(url)
}

// Implementing from root package
func (u *UseInterface) UseHuhMenuStart() {
	HuhMenuStart()
}
func (u *UseInterface) UseQuitWithSpinner() {
	QuitWithSpinner()
}

// Using SemChoose from semChoose package
func (u *UseInterface) UseSemChoose(url string) {
	semChoose.SemChoose(url)
}

// Using SemTable from semTable package
func (u *UseInterface) UseSemTable(url string) {
	semTable.SemTable(url)
}

// Using Year from year package
func (u *UseInterface) UseYear(url string) {
	year.Year(url)
}

type Subject struct {
	name string
	path string
}

var RootCmd = &cobra.Command{
	Use:   "ampyq",
	Short: "Amrita PYQ CLI",
	Long:  `A CLI application to access Amrita Repository for previous year question papers.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(helpers.LogoStyle.Render(logo.LOGO_ASCII))
		HuhMenuStart()
	},
}

func QuitWithSpinner() {
	action := func() {
		time.Sleep(2 * time.Second)
	}

	if err := spinner.New().
		Type(spinner.Line).
		Title("  Exiting ...").
		TitleStyle(helpers.FetchStatusStyle.Inline(true)).
		Action(action).
		Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}

func HuhMenuStart() {
	action := func() {
		time.Sleep(2 * time.Second)
	}
	if err := spinner.New().Title("Fetching Courses").Action(action).Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	resources, err := requestClient.GetCoursesReq(configs.COURSE_LIST_URL)
	if err != nil {
		fmt.Println(helpers.ErrorStyle.Render(fmt.Sprintf("Error: %v\n", err)))
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

	// Auto-exit if "Quit" is selected
	if selectedOption == "Quit" {
		QuitWithSpinner()
	}

	// Store only the selected course
	helpers.SelectionHistory = []string{selectedOption} // Reset history to show only last selected

	// Move to the next menu (Second Menu)
	for _, subject := range subjects {
		if subject.name == selectedOption {
			url := configs.BASE_URL + subject.path
			semTable.SemTable(url) // Show only selected course in second menu
			break
		}
	}
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
