package main

import (
	"bytes"
	"flag"
	"fmt"
	"math"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type Cycle struct {
	issue    string
	end      time.Time
	duration time.Duration
}

func printCycleTimes(path string, authorExclude regexp.Regexp, days int, github bool) {
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

	cycles := make(map[string]Cycle)

	issueRegex, _ := regexp.Compile("#([0-9]+)")

	commits.ForEach(func(c *object.Commit) error {
		issue := issueRegex.Find([]byte(c.Message))
		if issue == nil {
			return nil
		}

		if authorExclude.Match([]byte(c.Author.Name)) {
			return nil
		}

		if days > 0 {
			timeOfCommit := c.Author.When
			deltaDuration := time.Duration(days) * 24 * time.Hour
			earliestOfInterest := time.Now().Add(-deltaDuration)
			if timeOfCommit.Before(earliestOfInterest) {
				return nil
			}
		}

		existingCycle, cycleKnown := cycles[string(issue)]
		var end time.Time
		if cycleKnown {
			end = existingCycle.end
		} else {
			end = c.Committer.When
		}

		cycles[string(issue)] = Cycle{
			issue:    string(issue),
			end:      end,
			duration: end.Sub(c.Committer.When),
		}

		return nil
	})

	result := make([]Cycle, 0, len(cycles))

	for _, value := range cycles {
		result = append(result, value)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].end.Before(result[j].end)
	})

	for _, cycle := range result {
		visualRepresentation := strings.Repeat("·", int(math.Ceil(cycle.duration.Hours()/8)))
		if len(visualRepresentation) > 50 {
			visualRepresentation = strings.Repeat("·", 48) + ">>"
		}
		issue := cycle.issue[1:]

		if github {
			ghCmd := exec.Command("gh")

			ghCmd.Dir = path
			var outb, errb bytes.Buffer
			ghCmd.Stdout = &outb
			ghCmd.Stderr = &errb

			ghCmd.Args = []string{"gh", "issue", "view", "--json", "title", "-t", "{{.title}}", issue}

			err := ghCmd.Run()

			if err != nil {
				fmt.Println(err)
				fmt.Println(errb.String())
			} else {
				title := outb.String()
				if len(title) > 60 {
					title = fmt.Sprintf("%s...", title[:53])
				}
				issue = fmt.Sprintf("%-60s #%s", title, issue)
			}
		} else {
			issue = fmt.Sprintf("#%s", issue)
		}

		fmt.Printf("%s %s\t%8.1f %s\n", cycle.end.Format(time.DateOnly), issue, cycle.duration.Hours(), visualRepresentation)
	}
}

func main() {
	excludeFlag := flag.String("exclude", "^$", "Exclude commits with authors that match this regex")
	daysFlag := flag.Int("days", -1, "How many days to look back, -1 being infinity")
	githubFlag := flag.Bool("gh", false, "Use gh cli to obtain issue titles")

	flag.Usage = func() {
		fmt.Print("Usage: cycletime [flags] [PATH]\n\n")
		fmt.Print("Hours between first and last commit tagged with an issue number\n\n")
		fmt.Print("PATH defaults to the current working directory\n\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	authorExcludeRegex, err := regexp.Compile(*excludeFlag)

	if err != nil {
		fmt.Println("Invalid regex")
		fmt.Println(err)
		flag.Usage()
		os.Exit(1)
	}

	path := flag.Arg(0)

	if path == "" {
		path, err = os.Getwd()
		if err != nil {
			fmt.Println("Can't get current working directory")
			os.Exit(1)
		}
	}

	printCycleTimes(path, *authorExcludeRegex, *daysFlag, *githubFlag)
}
