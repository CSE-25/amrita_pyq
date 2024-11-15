package cmd

import (
	"fmt"
	"os"
)

func year(url string) {
	fmt.Println("Fetching...")
	year_url := url

	files, err := yearReq(url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("No\tFiles\n")
	for i, file := range files {
		fmt.Printf("%d\t%s\n", i+1, file.name)
	}

	// Option to add "Back".
	fmt.Printf("%d\tBack\n", len(files)+1)

	for {
		var ch int
		fmt.Print("\nEnter your choice: ")
		fmt.Scanln(&ch)

		if ch > 0 && ch <= len(files) {
			link := files[ch-1].path
			url = BASE_URL + link
			break
		} else if ch == len(files)+1 {
			semChoose(stack.Pop())
		} else {
			fmt.Println("Please enter a valid input!")
		}
	}

	fmt.Println("Please wait until browser opens !")
	openBrowser(url)

	var ch int
	fmt.Println("Do you want to continue ? \nPress 1 for Yes and 0 for No : ")
	fmt.Scanln(&ch)

	if ch == 0 {
		fmt.Println("Exiting...")
		os.Exit(0)
	} else {
		year(year_url)
	}
}
