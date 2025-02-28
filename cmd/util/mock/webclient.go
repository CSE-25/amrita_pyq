package mock

import (
	"errors"
)

// MockWebClient is a mock implementation of the WebClient interface.
type MockWebClient struct {
	FetchHTMLFunc   func(url string) (string, error)
	OpenBrowserFunc func(url string) error
	DownloadFileFunc func(url, filename string) error
}

// FetchHTML mocks the FetchHTML method of the WebClient interface.
func (m MockWebClient) FetchHTML(url string) (string, error) {
	if m.FetchHTMLFunc != nil {
		return m.FetchHTMLFunc(url)
	}
	return "", errors.New("FetchHTML not implemented")
}

// OpenBrowser mocks the OpenBrowser method of the WebClient interface.
func (m MockWebClient) OpenBrowser(url string) error {
	if m.OpenBrowserFunc != nil {
		return m.OpenBrowserFunc(url)
	}
	return errors.New("OpenBrowser not implemented")
}

// DownloadFile mocks the DownloadFile method of the WebClient interface.
func (m MockWebClient) DownloadFile(url, filename string) error {
	if m.DownloadFileFunc != nil {
		return m.DownloadFileFunc(url, filename)
	}
	return errors.New("DownloadFile not implemented")
}