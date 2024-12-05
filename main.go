package main

// Issues prints a table of GitHub issues matching the search terms.
import (
	"fmt"
	"log"
	"os"

	"github.com/viiku/issue-tracker/pkg/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s %t \n",
			item.Number, item.User.Login, item.Title, item.IsLocked)
		// fmt.Println(item.User.GITHUBID)
		// fmt.Println(item.User.HTMLURL)
		// fmt.Println(item.User.FOLLOWERSURL)
	}
}
