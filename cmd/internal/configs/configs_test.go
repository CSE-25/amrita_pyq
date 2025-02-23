package configs

import (
	"testing"
)

func TestConfigs(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		wantUrl string
	}{
		{
			name:    "SuccessBaseURLDeclaredProperly",
			url:     BASE_URL,
			wantUrl: "http://dspace.amritanet.edu:8080",
		},
		{
			name:    "SuccessCourseURLDeclaredProperly",
			url:     COURSE_URL,
			wantUrl: "http://dspace.amritanet.edu:8080/xmlui/handle/123456789/",
		},
		{
			name:    "SuccessCourseListURLDeclaredProperly",
			url:     COURSE_LIST_URL,
			wantUrl: "http://dspace.amritanet.edu:8080/xmlui/handle/123456789/16",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if tc.url != tc.wantUrl {
				t.Errorf("Expected %s, got %s", tc.wantUrl, tc.url)
			}
		})
	}
}
