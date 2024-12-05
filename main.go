package main

// Issues prints a table of GitHub issues matching the search terms.
import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/viiku/issue-tracker/pkg/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	now := time.Now()

	var lessThanMonth, lessThanYear, moreThanYear []*github.Issue
	for _, item := range result.Items {
		// fmt.Printf("#%-5d %9.9s %.55s %t \n",
		// 	item.Number, item.User.Login, item.Title, item.IsLocked)
		// fmt.Println(item.User.GITHUBID)
		// fmt.Println(item.User.HTMLURL)
		// fmt.Println(item.User.FOLLOWERSURL)

		// fmt.Printf("%d issues:\n", result.TotalCount)

		age := now.Sub(item.CreatedAt)

		switch {
		case age < 30*24*time.Hour:
			lessThanMonth = append(lessThanMonth, item)
		case age < 365*24*time.Hour:
			lessThanYear = append(lessThanYear, item)
		default:
			moreThanYear = append(moreThanYear, item)
		}
	}

	// Display issues in each category
	fmt.Println("\nIssues less than a month old:")
	for _, item := range lessThanMonth {
		fmt.Printf("#%-5d %9.9s %.55s %s\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt.Format("2006-01-02 15:04:05"))
	}

	fmt.Println("\nIssues less than a Year old:")
	for _, item := range lessThanYear {
		fmt.Printf("#%-5d %9.9s %.55s %s\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt.Format("2006-01-02 15:04:05"))
	}

	fmt.Println("\nIssues more than a Year old:")
	for _, item := range moreThanYear {
		fmt.Printf("#%-5d %9.9s %.55s %s\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt.Format("2006-01-02 15:04:05"))
	}
}
