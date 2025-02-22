package helpers

import (
	"testing"
)

func TestFetchHTML(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		wantErr bool
	}{
		{
			name:    "SuccessValidURLFetch",
			url:     "https://httpbin.org/get",
			wantErr: false,
		},
		{
			name:    "FailInvalidURLFetch",
			url:     "http://invalid.invalid",
			wantErr: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			_, gotErr := FetchHTML(tc.url)
			if tc.wantErr && gotErr == nil {
				t.Errorf("expected error for URL %s, got nil", tc.url)
			} else if !tc.wantErr && gotErr != nil {
				t.Errorf("unexpected error for URL %s: %v", tc.url, gotErr)
			}
		})
	}
}

func TestOpenBrowser(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping TestOpenBrowser in short mode")
	}

	tests := []struct {
		name    string
		url     string
		wantErr bool
	}{
		{
			name:    "SuccessOpenValidURLInBrowser",
			url:     "https://httpbin.org/get",
			wantErr: false,
		},
		// TODO: Add test cases to cover error scenarios.
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if gotErr := OpenBrowser(tc.url); (gotErr != nil) != tc.wantErr {
				t.Errorf("for URL %s, expected error: %v, got: %v", tc.url, tc.wantErr, gotErr)
			}
		})
	}
}
