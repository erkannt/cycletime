package main

import (
	"fmt"
	"os"

	"gopkg.in/src-d/go-git.v4"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println("Failed to get current working directory")
		os.Exit(1)
	}

	repo, err := git.PlainOpen(path)
	if err != nil {
		fmt.Println("Could not read git repo in current working directory")
		fmt.Println(err)
		os.Exit(1)
	}

	repo.Log(&git.LogOptions{})

	fmt.Println(path)
}
