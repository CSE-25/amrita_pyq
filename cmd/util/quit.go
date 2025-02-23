package util

import (
	"fmt"
	"os"
	"time"

	"amrita_pyq/cmd/internal/configs"
	"github.com/charmbracelet/huh/spinner"
)

// QuitWithSpinner quits the TUI
// with a spinner animation.
func QuitWithSpinner() {
	action := func() {
		time.Sleep(2 * time.Second)
	}

	if err := spinner.New().
		Type(spinner.Line).
		Title("  Exiting ...").
		TitleStyle(configs.FetchStatusStyle.Inline(true)).
		Action(action).
		Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}
