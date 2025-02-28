package webclient

import (
	"amrita_pyq/cmd/internal/configs"
	"fmt"
	"net/http"
	"os"
	"io"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/anaskhan96/soup"
)

// WebClient defines methods for fetching HTML and opening a browser.
type WebClient interface {
	FetchHTML(url string) (string, error)
	OpenBrowser(url string) error
	DownloadFile(url, filename string) error
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

// DownloadFile downloads a file from the given URL and saves it to the pyq_downloads folder.
func (d DefaultWebClient) DownloadFile(url, filename string) error {
	downloadFolder := "pyq_downloads"
	if err := os.MkdirAll(downloadFolder, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create download folder: %v", err)
	}

	filePath := filepath.Join(downloadFolder, filename)
	out, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download file: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned non-200 status: %d", resp.StatusCode)
	}
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("error writing file: %v", err)
	}

	return nil
}

