package interfaces

import (
	"amrita_pyq/cmd/model"

	"github.com/charmbracelet/lipgloss"
)

type Interface interface {

	//configs
	UseBASE_URL() string
	UseCOURSE_LIST_URL() string

	//helpers
	UseLogoStyle() lipgloss.Style
	UseErrorStyle() lipgloss.Style
	UseFetchStatusStyle() lipgloss.Style
	UseFetchHTML(url string) (string, error)
	UseOpenBrowser(url string) error

	//logo
	UseLOGO_ASCII() string

	//requestClient
	UseGetCoursesReq(url string) ([]model.Resource, error)
	UseSemChooseReq(url string) ([]model.Resource, error)
	UseSemTableReq(url string) ([]model.Resource, error)
	UseYearReq(url string) ([]model.Resource, error)

	//root
	UseHuhMenuStart()
	UseQuitWithSpinner()

	//semChoose
	UseSemChoose(url string)

	//semTable
	UseSemTable(url string)

	//year
	UseYear(url string)
}
