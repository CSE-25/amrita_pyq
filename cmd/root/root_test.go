package root

import (
	"amrita_pyq/cmd/model"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestRoot(t *testing.T) {

	//Test RootCmd
	if _, ok := interface{}(RootCmd).(*cobra.Command); !ok {
		t.Errorf("RootCmd is not of type *cobra.Command")
	}

	if RootCmd.Use != "ampyq" {
		t.Errorf("Expected Use: 'ampyq', Got: %v", RootCmd.Use)
	}

	if RootCmd.Short != "Amrita PYQ CLI" {
		t.Errorf("Expected Short: 'Amrita PYQ CLI', Got: %v", RootCmd.Short)
	}

	long := `A CLI application to access Amrita Repository for previous year question papers.`
	if RootCmd.Long != long {
		t.Errorf("Expected Long: %q, Got: %v", long, RootCmd.Long)
	}
}

type SampleRequestClient struct{}

func (m *SampleRequestClient) GetCoursesReq(url string) ([]model.Resource, error) {
	return []model.Resource{
		{Name: "Course 1", Path: "/course1"},
		{Name: "Course 2", Path: "/course2"},
	}, nil
}

func TestHuhMenuStart(t *testing.T) {
	mockClient := &SampleRequestClient{}

	resources, err := mockClient.GetCoursesReq("https://httpbin.org/get")

	// Test the output
	assert.Nil(t, err)
	assert.Len(t, resources, 2)
	assert.Equal(t, "Course 1", resources[0].Name)
	assert.Equal(t, "/course1", resources[0].Path)
}
