package webclient

import (
	"amrita_pyq/cmd/internal/configs"
	"fmt"
	"os/exec"
	"runtime"

	"github.com/anaskhan96/soup"
)

// WebClient defines methods for fetching HTML and opening a browser.
type WebClient interface {
	FetchHTML(url string) (string, error)
	OpenBrowser(url string) error
}

// DefaultWebClient implements WebClient using real network calls.
type DefaultWebClient struct{}

// FetchHTML fetches and parses HTML from the given URL.
func (d DefaultWebClient) FetchHTML(url string) (string, error) {
	doc, err := soup.Get(url)
	if err != nil {
		fmt.Println(configs.ErrorStyle.Render("Error fetching the URL. Make sure you're connected to Amrita WiFi or VPN."))
		return "", err
	}
	return doc, nil
}

// OpenBrowser opens a URL in the default web browser.
func (d DefaultWebClient) OpenBrowser(url string) error {
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
		styledMessage := configs.ErrorStyle.Render("failed to open browser")
		return fmt.Errorf("%s: %w", styledMessage, err)
	}

	return nil
}
