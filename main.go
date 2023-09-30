package main

import (
	"fmt"
	"os"

	"github.com/sosukesuzuki/gh-sosuke/cmd"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("GitHub CLI extension for sosukesuzuki's OSS development.")
		os.Exit(0)
	}
	switch os.Args[1] {
	case "issue":
		cmd.Issue()
	default:
		fmt.Fprintf(os.Stderr, "[ERROR] gh sosuke doesn't support `%s`\n", os.Args[1])
		os.Exit(1)
	}
}

// For more examples of using go-gh, see:
// https://github.com/cli/go-gh/blob/trunk/example_gh_test.go
