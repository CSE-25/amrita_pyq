package helpers

import (
	"testing"
)

func TestFetchHTML(t *testing.T) {
	_, err := FetchHTML("https://httpbin.org/get")
	if err != nil {
		t.Errorf("Failed to fetch HTML content: %v", err)
	}
}

func TestOpenBrowser(t *testing.T) {
	err := OpenBrowser("https://httpbin.org/get")
	if err != nil {
		t.Errorf("Failed to open browser: %v", err)
	}
}
