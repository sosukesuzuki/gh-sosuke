package cmd

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	"github.com/cli/go-gh/v2"
	"github.com/cli/go-gh/v2/pkg/api"
)

func calculateHoursAgo(timeStr string) (string, error) {
	parsedTime, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return "", err
	}

	currentTime := time.Now()
	duration := currentTime.Sub(parsedTime)

	hours := math.Floor(duration.Hours() + 0.5)

	return fmt.Sprintf("%.0f", hours), nil
}

func notificationList(client *api.RESTClient) (string, error) {
	nameWithOwner, _, err := gh.Exec("repo", "view", "--json", "nameWithOwner", "-q", ".nameWithOwner")
	if err != nil {
		return "", err
	}

	apiPath := fmt.Sprintf("repos/%s/notifications", strings.TrimSuffix(nameWithOwner.String(), "\n"))

	response := []struct {
		Subject struct {
			Title string
			Url   string
			Type  string
		}
		Updated_at string
	}{}

	err = client.Get(apiPath, &response)
	if err != nil {
		return "", err
	}
	if len(response) == 0 {
		return "", nil
	}

	notifications := make([]string, len(response))
	for i, notification := range response {
		number := path.Base(notification.Subject.Url)
		title := notification.Subject.Title
		notifType := notification.Subject.Type
		hoursAgo, err := calculateHoursAgo(notification.Updated_at)

		if err != nil {
			return "", err
		}

		notifications[i] = fmt.Sprintf("%s [%s] %s  %s hours ago", number, notifType, title, hoursAgo)
	}
	notifList := strings.Join(notifications, "\n")

	pecoCmd := exec.Command("peco")
	pecoCmd.Stdin = strings.NewReader(notifList)
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

func Notification() error {
	client, err := api.DefaultRESTClient()

	if err != nil {
		return err
	}

	switch os.Args[2] {
	case "list":
		out, err := notificationList(client)
		if err != nil {
			return err
		}
		fmt.Fprintf(os.Stdout, "%s", out)
		return nil
	default:
		return fmt.Errorf("gh sosuke doesn't support `%s`", os.Args[2])
	}
}
