package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type GhSosukeIssueError struct {
	msg string
	err error
}

func (e *GhSosukeIssueError) Error() string {
	return fmt.Sprintf("Error from `gh sosuke issue`: %s\n%s", e.msg, e.err.Error())
}
func (e *GhSosukeIssueError) Unwrap() error {
	return e.err
}

func list() (string, error) {
	args := os.Args[3:]

	cmdStr := fmt.Sprintf("gh issue list %s | peco", strings.Join(args, " "))
	cmd := exec.Command("bash", "-c", cmdStr)
	out, err := cmd.Output()
	if err != nil {
		return "", &GhSosukeIssueError{msg: "Failed to run `gh issue list`", err: err}
	}
	outStr := string(out)
	return strings.Fields(outStr)[0], nil
}

func targetedCommand() (string, error) {
	if len(os.Args) == 3 {
		return "", &GhSosukeIssueError{msg: "`gh sosuke issue` needs issue number"}
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputValue := scanner.Text()
	if err := scanner.Err(); err != nil {
		return "", &GhSosukeIssueError{msg: "Failed to read input", err: err}
	}

	cmdStr := fmt.Sprintf("gh issue %s %s", strings.Join(os.Args[3:], " "), inputValue)

	cmd := exec.Command("bash", "-c", cmdStr)
	out, err := cmd.Output()
	if err != nil {
		return "", &GhSosukeIssueError{msg: "Failed to run `gh issue`", err: err}
	}
	outStr := string(out)
	return outStr, nil
}

func Issue() error {
	if len(os.Args) == 2 {
		return &GhSosukeIssueError{msg: fmt.Sprintf("`gh sosuke issue` require sub-command")}
	}

	switch os.Args[2] {
	case "list":
		out, err := list()
		if err != nil {
			return err
		}
		fmt.Fprintf(os.Stdout, "%s", out)
		return nil
	case "close", "create", "delete", "develop", "edit", "lock", "pin", "reopen", "transfer", "unlock", "unpin", "view":
		out, err := targetedCommand()
		if err != nil {
			return err
		}
		fmt.Fprintf(os.Stdout, "%s", out)
		return nil
	default:
		return &GhSosukeIssueError{msg: fmt.Sprintf("`gh sosuke issue` doesn't support `%s`", os.Args[2])}
	}
}
