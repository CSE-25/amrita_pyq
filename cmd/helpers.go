package cmd

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"log"
	"os/exec"
	"runtime"
)

// Fetches and parses HTML from the given URL.
func fetchHTML(url string) (string, error) {
	doc, err := soup.Get(url)

	if err != nil {
		fmt.Println("Error fetching the URL. Make sure you're connected to Amrita WiFi or VPN.")
		return "", err
	}

	return doc, nil
}

// Opens a URL in the default web browser.
func openBrowser(url string) {
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
		log.Fatal(err)
	}
}
