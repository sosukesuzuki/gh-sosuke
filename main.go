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

	var err error
	switch os.Args[1] {
	case "issue":
		err = cmd.Issue()
	case "notification", "notif":
		err = cmd.Notification()
	default:
		err = fmt.Errorf("gh sosuke doesn't support `%s`", os.Args[1])
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] %s\n", err.Error())
		os.Exit(1)
	}
}

// For more examples of using go-gh, see:
// https://github.com/cli/go-gh/blob/trunk/example_gh_test.go
