package version

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {

	//Test versionCmd
	if _, ok := interface{}(versionCmd).(*cobra.Command); !ok {
		t.Errorf("versionCmd is not of type *cobra.Command")
	}
	if versionCmd.Use != "version" {
		t.Errorf("Expected Use to be 'version', got %v", versionCmd.Use)
	}
	if versionCmd.Short != "Print the version number of ampyq" {
		t.Errorf("Unexpected Short description: %v", versionCmd.Short)
	}
	if versionCmd.Long != "Displays version of ampyq installed on the system." {
		t.Errorf("Unexpected Long description: %v", versionCmd.Long)
	}

	var out bytes.Buffer
	versionCmd.SetOut(&out)

	versionCmd.Run(versionCmd, []string{})

	// Check the output
	expectedOutput := "Amrita Previous Year Questions v0.0.1-alpha\n"
	assert.Equal(t, expectedOutput, out.String())
}
