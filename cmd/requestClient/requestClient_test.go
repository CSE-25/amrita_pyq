package requestClient

import (
	"testing"
)

func TestGetCoursesReq(t *testing.T) {

	//Testing GetCoursesReq
	_, err := GetCoursesReq("testing_for_error")

	if err == nil {
		t.Errorf("Expected: Error, Got: nil")
	}

	if err != errHTMLFetch {
		t.Errorf("Expected: %v, Got: %v", errHTMLFetch, err)
	}
}

func TestSemChooseReq(t *testing.T) {

	//Testing SemChooseReq
	_, err := SemChooseReq("testing_for_error")

	if err == nil {
		t.Errorf("Expected: Error, Got: nil")
	}

	if err != errHTMLFetch {
		t.Errorf("Expected: %v, Got: %v", errHTMLFetch, err)
	}
}

func TestSemTableReq(t *testing.T) {

	//Testing SemTableReq
	_, err := SemTableReq("testing_for_error")

	if err == nil {
		t.Errorf("Expected: Error, Got: nil")
	}

	if err != errHTMLFetch {
		t.Errorf("Expected: %v, Got: %v", errHTMLFetch, err)
	}
}

func TestYearReq(t *testing.T) {

	//Testing YearReq
	_, err := YearReq("testing_for_error")

	if err == nil {
		t.Errorf("Expected: Error, Got: nil")
	}

	if err != errHTMLFetch {
		t.Errorf("Expected: %v, Got: %v", errHTMLFetch, err)
	}
}
