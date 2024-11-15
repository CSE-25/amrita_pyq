package cmd

import (
	"fmt"
)

func semChoose(url string) {
	fmt.Println("Fetching assesments...")
	params_url := url

	assesments, err := semChooseReq(url)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Display the found items
	fmt.Printf("No\tSemesters\n")
	for i, assesment := range assesments {
		fmt.Printf("%d\t%s\n", i+1, assesment.name) // Extract the text from the span element
	}

	// Option to add "Back"
	fmt.Printf("%d\tBack\n", len(assesments)+1)

	for {
		var ch int
		fmt.Print("\nEnter your assesment: ")
		fmt.Scanln(&ch)

		if ch > 0 && ch <= len(assesments) {
			url = BASE_URL + assesments[ch-1].path
			break
		} else if ch == len(assesments)+1 {
			semTable(stack.Pop())
		} else {
			fmt.Println("Please enter a valid input!")
		}
	}

	// append to stack
	stack.Push(params_url)

	year(url)
}
