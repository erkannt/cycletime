package main

import (
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type cycle struct {
	issue    string
	end      time.Time
	duration time.Duration
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("No PATH provided")
		fmt.Println("cycletime PATH")
		os.Exit(1)
	}
	path := os.Args[1]

	repo, err := git.PlainOpen(path)
	if err != nil {
		fmt.Printf("Not a git repo: %s\n", path)
		os.Exit(1)
	}

	commits, err := repo.Log(&git.LogOptions{Order: git.LogOrderCommitterTime})
	if err != nil {
		fmt.Printf("Failed to load git log:\n%s\n", err)
		os.Exit(1)
	}

	cycles := make(map[string]cycle)

	issueRegex, _ := regexp.Compile("#([0-9]+)")

	commits.ForEach(func(c *object.Commit) error {
		issue := issueRegex.Find([]byte(c.Message))
		if issue == nil {
			return nil
		}

		existingCycle, cycleKnown := cycles[string(issue)]
		var end time.Time
		if cycleKnown {
			end = existingCycle.end
		} else {
			end = c.Committer.When
		}

		cycles[string(issue)] = cycle{
			issue:    string(issue),
			end:      end,
			duration: end.Sub(c.Committer.When),
		}

		newCycle := cycles[string(issue)]
		fmt.Printf("%s %.0fh\n", newCycle.issue, newCycle.duration.Hours())

		return nil
	})

}
