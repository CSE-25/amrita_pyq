package webclient

import (
	"amrita_pyq/cmd/internal/configs"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"github.com/anaskhan96/soup"
)

// WebClient defines methods for fetching HTML, opening a browser, and downloading files.
type WebClient interface {
	FetchHTML(url string) (string, error)
	OpenBrowser(url string) error
	DownloadFile(url string, filename string) error
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

// DownloadFile downloads a file from the URL and saves it to the local disk.
func (d DefaultWebClient) DownloadFile(url string, filename string) error {
	// Create the file
	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
