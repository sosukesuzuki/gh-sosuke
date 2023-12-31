package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/cli/go-gh/v2"
)

func issueList() (string, error) {
	ghCmdStr := append([]string{"issue", "list"}, os.Args[3:]...)
	issueList, _, err := gh.Exec(ghCmdStr...)

	if err != nil {
		return "", err
	}

	pecoCmd := exec.Command("peco")
	pecoCmd.Stdin = strings.NewReader(issueList.String())
	pecoCmd.Stderr = os.Stderr
	var pecoStdout strings.Builder
	pecoCmd.Stdout = &pecoStdout

	err = pecoCmd.Run()
	if err != nil {
		return "", err
	}

	fields := strings.Fields(pecoStdout.String())

	if len(fields) == 0 {
		return "", nil
	}

	return fields[0], nil
}

func issueTargetedCommand() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputValue := scanner.Text()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	ghCmdStr := append([]string{"issue", os.Args[2], inputValue}, os.Args[3:]...)
	ghCmdResult, _, err := gh.Exec(ghCmdStr...)

	if err != nil {
		return "", err
	}
	outStr := ghCmdResult.String()

	return outStr, nil
}

func Issue() error {
	switch os.Args[2] {
	case "list":
		out, err := issueList()
		if err != nil {
			return err
		}
		fmt.Fprintf(os.Stdout, "%s", out)
		return nil
	case "close", "create", "delete", "develop", "edit", "lock", "pin", "reopen", "transfer", "unlock", "unpin", "view":
		out, err := issueTargetedCommand()
		if err != nil {
			return err
		}
		fmt.Fprintf(os.Stdout, "%s", out)
		return nil
	default:
		return fmt.Errorf("`gh sosuke issue` doesn't support `%s`", os.Args[2])
	}
}
