package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

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

	commits, err := repo.Log(&git.LogOptions{})
	if err != nil {
		fmt.Printf("Failed to load git log:\n%s\n", err)
	}

	var cCount int
	commits.ForEach(func(c *object.Commit) error {
		cCount++
		return nil
	})

	fmt.Println("Commit count:", cCount)
}
