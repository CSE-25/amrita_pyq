package model

import (
	"testing"
)

func TestModel(t *testing.T) {

	//Testing Resource
	expResource := Resource{"Testing", "Testing/testing"}
	if expResource.Name != "Testing" || expResource.Path != "Testing/testing" {
		t.Errorf("Expected %v, Received %v", expResource, Resource{"Testing", "Testing/testing"})
	}
}
