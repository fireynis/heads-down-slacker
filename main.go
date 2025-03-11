package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/slack-go/slack"
)

type config struct {
	token       string
	channels    []string
	duration    time.Duration
	emoji       string
	statusText  string
	message     string
	backMessage string
	headingDown bool
}

func main() {
	// Parse command line flags
	headingDown := flag.Bool("down", true, "Set to true to go heads down, false to revert")
	token := flag.String("token", "", "Slack API token")
	channels := flag.String("channels", "", "Comma-separated list of channel IDs to post to")
	duration := flag.Duration("duration", 1*time.Hour, "Duration to be heads down (e.g. 1h, 30m)")
	emoji := flag.String("emoji", ":no_entry:", "Emoji for the status")
	statusText := flag.String("status", "Heads Down - Focus Time", "Status text")
	message := flag.String("message", "I'll be heads down for %s. Will respond afterward.", "Message to post (use %s for duration)")
	backMessage := flag.String("backmessage", "I'm back online and available now.", "Message to post when coming back online")
	flag.Parse()

	// Check for token in env var if not provided in flag
	if *token == "" {
		*token = os.Getenv("SLACK_TOKEN")
		if *token == "" {
			log.Fatal("Slack token required: provide with -token flag or SLACK_TOKEN environment variable")
		}
	}

	config := config{
		token:       *token,
		headingDown: *headingDown,
		emoji:       *emoji,
		statusText:  *statusText,
		duration:    *duration,
		message:     *message,
		backMessage: *backMessage,
	}

	if *channels != "" {
		config.channels = strings.Split(*channels, ",")
	}

	// Create Slack client
	api := slack.New(config.token)

	// Handle heads down or heads up
	if config.headingDown {
		err := goHeadsDown(api, config)
		if err != nil {
			log.Fatalf("Error going heads down: %v", err)
		}
	} else {
		err := revertHeadsDown(api, config)
		if err != nil {
			log.Fatalf("Error reverting heads down status: %v", err)
		}
	}
}

func goHeadsDown(api *slack.Client, config config) error {
	// Set user status
	err := setStatus(api, config.emoji, config.statusText)
	if err != nil {
		return fmt.Errorf("failed to set status: %w", err)
	}
	fmt.Println("Status updated to heads down")

	// Post to channels if provided
	if len(config.channels) > 0 {
		durationStr := formatDuration(config.duration)
		message := fmt.Sprintf(config.message, durationStr)

		for _, channelID := range config.channels {
			err := sendMessage(api, channelID, message)
			if err != nil {
				fmt.Printf("Warning: Failed to post to channel %s: %v\n", channelID, err)
				continue
			}
			fmt.Printf("Posted heads down message to channel %s\n", channelID)
		}
	}

	return nil
}

func revertHeadsDown(api *slack.Client, config config) error {
	// Clear user status
	err := setStatus(api, "", "")
	if err != nil {
		return fmt.Errorf("failed to clear status: %w", err)
	}
	fmt.Println("Status cleared")

	// Post "back online" message to channels if provided
	if len(config.channels) > 0 {
		for _, channelID := range config.channels {
			err := sendMessage(api, channelID, config.backMessage)
			if err != nil {
				fmt.Printf("Warning: Failed to post to channel %s: %v\n", channelID, err)
				continue
			}
			fmt.Printf("Posted back online message to channel %s\n", channelID)
		}
	}

	return nil
}

func sendMessage(api *slack.Client, channelID string, message string) error {
	_, _, err := api.PostMessage(channelID, slack.MsgOptionText(message, false))
	return err
}

func setStatus(api *slack.Client, emoji string, text string) error {
	// Set user profile
	profile := slack.UserProfile{
		StatusText:  text,
		StatusEmoji: emoji,
	}
	err := api.SetUserCustomStatus(profile.StatusText, profile.StatusEmoji, 0)
	return err
}

func formatDuration(d time.Duration) string {
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute

	if h > 0 && m > 0 {
		return fmt.Sprintf("%dh %dm", h, m)
	} else if h > 0 {
		return fmt.Sprintf("%dh", h)
	}
	return fmt.Sprintf("%dm", m)
}
