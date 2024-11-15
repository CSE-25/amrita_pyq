package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ampyq",
	Short: "Amrita PYQ CLI",
	Long:  `A CLI application to access Amrita Repository for previous year question papers.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(LOGO_ASCII)
		start()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// start function - equivalent to start() in Python
func start() {
	fmt.Println("Fetching Courses...")
	subjects, err := getCoursesReq(COURSE_LIST_URL)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Available Courses:")

	for i, subject := range subjects {
		fmt.Printf("%d.\t%s\n", i+1, subject.name)
	}

	// Option to quit.
	fmt.Printf("%d.\tQuit\n", len(subjects)+1)

	for {
		var ch int
		fmt.Printf("\nEnter your choice: ")
		fmt.Scanln(&ch)

		if ch > 0 && ch <= len(subjects) {
			path := subjects[ch-1].path
			url := BASE_URL + path
			semTable(url)
		} else if ch == len(subjects)+1 {
			fmt.Println("Goodbye!")
			os.Exit(0)
		} else {
			fmt.Println("Please enter a valid input!")
		}
	}
}
