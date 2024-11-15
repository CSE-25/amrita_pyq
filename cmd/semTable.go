package cmd

import (
	"fmt"
)

func semTable(url string) {
	fmt.Println("Fetching semesters...")

	semesters, err := semTableReq(url)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("No\tSemesters\n")
	for i, semester := range semesters {
		fmt.Printf("%d\t%s\n", i+1, semester.name)
	}

	// Option to add "Back".
	fmt.Printf("%d\tBack\n", len(semesters)+1)

	stack.Push(url)

	for {
		var ch int
		fmt.Print("\nEnter your semester: ")
		fmt.Scanln(&ch)

		if ch > 0 && ch <= len(semesters) {
			url := BASE_URL + semesters[ch-1].path
			semChoose(url)
			break
		} else if ch == len(semesters)+1 {
			start()
			break
		} else {
			fmt.Println("Please enter a valid input!")
		}
	}

}
