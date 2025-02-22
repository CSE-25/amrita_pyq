package helpers

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/anaskhan96/soup"
	"github.com/charmbracelet/lipgloss"
)

// Styling for the logo and error messages.
var (
	LogoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#01FAC6")).
			Bold(true)

	ErrorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("9")).
			Bold(true).
			Underline(true).
			Padding(0, 1).
			Margin(1, 0, 1, 0).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("1"))

	FetchStatusStyle = lipgloss.NewStyle().
				PaddingLeft(2).
				Foreground(lipgloss.Color("6")).
				Bold(true).
				Margin(1, 0)
)

var SelectionHistory []string

// Fetches and parses HTML from the given URL.
func FetchHTML(url string) (string, error) {
	doc, err := soup.Get(url)

	if err != nil {
		fmt.Println(ErrorStyle.Render("Error fetching the URL. Make sure you're connected to Amrita WiFi or VPN."))
		return "", err
	}

	return doc, nil
}

// Opens a URL in the default web browser.
func OpenBrowser(url string) error {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	}

	if err != nil {
		styledMessage := ErrorStyle.Render("failed to open browser")
		return fmt.Errorf("%s: %w", styledMessage, err)
	}

	return nil
}
