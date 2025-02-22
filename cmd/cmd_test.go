package cmd

import (
	"github.com/spf13/cobra"
	"testing"
)

func TestRoot(t *testing.T) {

	//Test RootCmd
	if _, ok := interface{}(Cmd).(*cobra.Command); !ok {
		t.Errorf("RootCmd is not of type *cobra.Command")
	}

	if Cmd.Use != "ampyq" {
		t.Errorf("Expected Use: 'ampyq', Got: %v", Cmd.Use)
	}

	if Cmd.Short != "Amrita PYQ CLI" {
		t.Errorf("Expected Short: 'Amrita PYQ CLI', Got: %v", Cmd.Short)
	}

	long := `A CLI application to access Amrita Repository for previous year question papers.`
	if Cmd.Long != long {
		t.Errorf("Expected Long: %q, Got: %v", long, Cmd.Long)
	}
}
