package configs

import (
	"testing"
)

func TestConfigs(t *testing.T) {

	//Testing base url
	expBASE_URL := "http://dspace.amritanet.edu:8080"
	if BASE_URL != expBASE_URL {
		t.Errorf("Expected %v, Received %v", expBASE_URL, BASE_URL)
	}

	//Testing course url
	expCOURSE_URL := expBASE_URL + "/xmlui/handle/123456789/"
	if COURSE_URL != expCOURSE_URL {
		t.Errorf("Expected %v, Received %v", expCOURSE_URL, COURSE_URL)
	}

	//Testing course list url
	expCOURSE_LIST_URL := expCOURSE_URL + "16"
	if COURSE_LIST_URL != expCOURSE_LIST_URL {
		t.Errorf("Expected %v, Received %v", expCOURSE_LIST_URL, COURSE_LIST_URL)
	}
}
