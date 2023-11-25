package chapter4

import (
	"fmt"
	"github"
	"log"
	"os"
)

func main() {
	input := os.Args[1:]
	fmt.Println(input)
	result, err := github.SearchIssues(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}

	fmt.Println("Sort issues by time:")
	sortLine := github.Parse(result)
	for key, val := range sortLine {
		fmt.Printf("%s: \n", key)
		for _, v := range val {
			fmt.Printf("#%-5d %9.9s %.55s %9.9v\n", v.Number, v.User.Login, v.Title, v.CreateAt)
		}
	}
}
