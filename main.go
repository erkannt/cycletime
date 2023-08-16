package main

import (
	"fmt"
	"os"

	"gopkg.in/src-d/go-git.v4"
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

	repo.Log(&git.LogOptions{})
}
