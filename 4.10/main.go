package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	"github.com/fredeastside/goex/4.10/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	sort.Sort(result)
	now := time.Now()
	monthText := "Last month\n"
	yearText := "Last year\n"
	moreThanYearText := "More than year ago\n"
	fmt.Printf("Themes: %d\n", result.TotalCount)
	for _, item := range result.Items {
		if item.CreatedAt.After(now.AddDate(0, -1, 0)) {
			fmt.Printf("%s", monthText)
			monthText = ""
		} else if item.CreatedAt.After(now.AddDate(-1, 0, 0)) {
			fmt.Printf("%s", yearText)
			yearText = ""
		} else if moreThanYearText != "" {
			fmt.Printf("%s", moreThanYearText)
			moreThanYearText = ""
		}
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
