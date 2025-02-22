package requestclient

import (
	"errors"
	"strings"
	"testing"

	"amrita_pyq/cmd/util/mock"
)

func newReqClient(fetchFunc func(string) (string, error)) RequestClient {
	return RequestClient{
		WebClient: mock.MockWebClient{FetchHTMLFunc: fetchFunc},
	}
}

func TestGetCoursesReq(t *testing.T) {
	successHTML := `<html>
		<body>
			<div id="aspect_artifactbrowser_CommunityViewer_div_community-view">
				<div class="artifact-title">
					<span>Course1</span>
					<a href="/course1"></a>
				</div>
			</div>
		</body>
	</html>`

	tests := []struct {
		name     string
		mockFunc func(url string) (string, error)
		inputURL string
		wantLen  int
		wantErr  bool
	}{
		{
			name: "SuccessCoursesHTML",
			mockFunc: func(url string) (string, error) {
				return successHTML, nil
			},
			inputURL: "/dummy",
			wantLen:  1,
			wantErr:  false,
		},
		{
			name: "FailHTMLError",
			mockFunc: func(url string) (string, error) {
				return "", errors.New("fetch error")
			},
			inputURL: "/dummy",
			wantLen:  0,
			wantErr:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			reqClient := newReqClient(tc.mockFunc)
			resources, err := reqClient.GetCoursesReq(tc.inputURL)
			if tc.wantErr && err == nil {
				t.Fatalf("expected error but got none")
			}
			if !tc.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(resources) != tc.wantLen {
				t.Fatalf("expected %d resources, got %d", tc.wantLen, len(resources))
			}
			if !tc.wantErr && tc.wantLen > 0 {
				if strings.TrimSpace(resources[0].Name) != "Course1" {
					t.Errorf("expected Course1 as resource name, got %s", resources[0].Name)
				}
				if resources[0].Path != "/course1" {
					t.Errorf("expected /course1 as resource path, got %s", resources[0].Path)
				}
			}
		})
	}
}

func TestSemChooseReq(t *testing.T) {
	successHTML := `<html>
		<body>
			<div id="aspect_artifactbrowser_CommunityViewer_div_community-view">
				<ul>
					<li><a href="/assessment1"><span>Assessment1</span></a></li>
					<li><a href="/assessment2"><span>Assessment2</span></a></li>
				</ul>
				<ul>
					<li><a href="/assessment3"><span>Assessment3</span></a></li>
				</ul>
			</div>
		</body>
	</html>`

	tests := []struct {
		name     string
		mockFunc func(url string) (string, error)
		inputURL string
		wantLen  int
		wantErr  bool
		expected Resource
	}{
		{
			name: "SuccessSemChooseHTML",
			mockFunc: func(url string) (string, error) {
				return successHTML, nil
			},
			inputURL: "/dummy",
			wantLen:  1,
			wantErr:  false,
			expected: Resource{Name: "Assessment3", Path: "/assessment3"},
		},
		{
			name: "FailHTMLError",
			mockFunc: func(url string) (string, error) {
				return "", errors.New("fetch error")
			},
			inputURL: "/dummy",
			wantLen:  0,
			wantErr:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			reqClient := newReqClient(tc.mockFunc)
			assessments, err := reqClient.SemChooseReq(tc.inputURL)
			if tc.wantErr && err == nil {
				t.Fatalf("expected error but got none")
			}
			if !tc.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(assessments) != tc.wantLen {
				t.Fatalf("expected %d assessments, got %d", tc.wantLen, len(assessments))
			}
			if !tc.wantErr && tc.wantLen > 0 {
				if assessments[0].Name != tc.expected.Name {
					t.Errorf("expected name %s, got %s", tc.expected.Name, assessments[0].Name)
				}
				if assessments[0].Path != tc.expected.Path {
					t.Errorf("expected path %s, got %s", tc.expected.Path, assessments[0].Path)
				}
			}
		})
	}
}

func TestSemTableReq(t *testing.T) {
	successHTML := `<html>
		<body>
			<div id="aspect_artifactbrowser_CommunityViewer_div_community-view">
				<ul>
					<li><a href="/sem1"><span>Semester1</span></a></li>
					<li><a href="/sem2"><span>Semester2</span></a></li>
				</ul>
			</div>
		</body>
	</html>`

	tests := []struct {
		name     string
		mockFunc func(url string) (string, error)
		inputURL string
		wantLen  int
		wantErr  bool
		expected []Resource
	}{
		{
			name: "SuccessSemTableHTML",
			mockFunc: func(url string) (string, error) {
				return successHTML, nil
			},
			inputURL: "/dummy",
			wantLen:  2,
			wantErr:  false,
			expected: []Resource{
				{Name: "Semester1", Path: "/sem1"},
				{Name: "Semester2", Path: "/sem2"},
			},
		},
		{
			name: "FailNoSemestersFound",
			mockFunc: func(url string) (string, error) {
				return `<html>
					<body>
						<div id="aspect_artifactbrowser_CommunityViewer_div_community-view">
							<ul></ul>
						</div>
					</body>
				</html>`, nil
			},
			inputURL: "/dummy",
			wantLen:  0,
			wantErr:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			reqClient := newReqClient(tc.mockFunc)
			semesters, err := reqClient.SemTableReq(tc.inputURL)
			if tc.wantErr && err == nil {
				t.Fatalf("expected error but got none")
			}
			if !tc.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !tc.wantErr && len(semesters) != tc.wantLen {
				t.Fatalf("expected %d semesters, got %d", tc.wantLen, len(semesters))
			}
			if !tc.wantErr && tc.wantLen > 0 {
				for i, expectedRes := range tc.expected {
					if semesters[i].Name != expectedRes.Name {
						t.Errorf("expected name %s, got %s", expectedRes.Name, semesters[i].Name)
					}
					if semesters[i].Path != expectedRes.Path {
						t.Errorf("expected path %s, got %s", expectedRes.Path, semesters[i].Path)
					}
				}
			}
		})
	}
}

func TestYearReq(t *testing.T) {
	firstHTML := `<html>
		<body>
			<div xmlns="http://di.tamu.edu/DRI/1.0/">
				<ul>
					<li><a href="/hyper"></a></li>
				</ul>
			</div>
		</body>
	</html>`

	secondHTML := `<html>
		<body>
			<div class="file-list">
				<div class="file-wrapper">
					<div><a href="/file1"></a></div>
					<div>
						<div>
							<span title="ignored"></span>
							<span title="File1"></span>
						</div>
					</div>
				</div>
			</div>
		</body>
	</html>`

	callCount := 0
	mockFunc := func(url string) (string, error) {
		callCount++
		switch callCount {
		case 1:
			return firstHTML, nil
		case 2:
			if !strings.Contains(url, "/hyper") {
				return "", errors.New("unexpected URL in second call")
			}
			return secondHTML, nil
		default:
			return "", errors.New("unexpected additional call")
		}
	}

	reqClient := newReqClient(mockFunc)
	resourceList, err := reqClient.YearReq("/dummy")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(resourceList) != 1 {
		t.Fatalf("expected 1 file resource, got %d", len(resourceList))
	}
	if resourceList[0].Name != "File1" {
		t.Errorf("expected resource name File1, got %s", resourceList[0].Name)
	}
	if resourceList[0].Path != "/file1" {
		t.Errorf("expected resource path /file1, got %s", resourceList[0].Path)
	}
}
