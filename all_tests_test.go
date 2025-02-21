package main

import (
	"amrita_pyq/cmd/configs"
	"os"
	"os/exec"
	"testing"
)

func TestAll(t *testing.T) {

	for _, path := range configs.TestPaths {
		t.Run(path, func(t *testing.T) {
			t.Helper()
			if err := runTestsInPackage(path); err != nil {
				t.Fatalf("Tests in %s failed: %v", path, err)
			}
		})
	}

}

func runTestsInPackage(pkg string) error {
	cmd := exec.Command("go", "test", "-v", "./"+pkg)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
