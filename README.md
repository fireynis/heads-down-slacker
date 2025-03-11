# Heads Down Slacker

A simple command-line tool to set your Slack status to "Heads Down" and notify channels that you'll be focusing.

## Installation

1. Build the binary:

```bash
go build -o heads-down-slacker
```

2. You can optionally copy the binary to a location in your PATH:

```bash
cp heads-down-slacker /usr/local/bin/
```

## Getting a Slack API Token

1. Go to [Slack API Apps](https://api.slack.com/apps)
2. Create a new app (or use an existing one)
3. Add the following OAuth scopes:
   - `users.profile:write` (to update your status)
   - `chat:write` (to post messages to channels)
4. Install the app to your workspace
5. Copy your OAuth Access Token

## Usage

### Setting Heads Down Status

```bash
# Basic usage (using SLACK_TOKEN environment variable)
export SLACK_TOKEN="xoxp-your-token-here"
./heads-down-slacker

# Specify token directly
./heads-down-slacker -token "xoxp-your-token-here"

# Customize duration (default: 1 hour)
./heads-down-slacker -duration 30m

# Customize emoji and status text
./heads-down-slacker -emoji ":focus:" -status "Deep work - please do not disturb"

# Notify specific channels (use channel IDs)
./heads-down-slacker -channels "C01234ABCDE,C056789FGHI"

# Custom message (use %s for duration placeholder)
./heads-down-slacker -message "I'm focusing for %s, will check messages after that."
```

### Reverting Status

```bash
./heads-down-slacker -down=false
```

## Getting Channel IDs

To find channel IDs, you can:
1. Open Slack in a browser
2. Navigate to the channel
3. The URL will contain the channel ID (e.g., `https://app.slack.com/client/TXXXXXXXX/CXXXXXXXX`)

## Creating a macOS shortcut or alias

Add this to your `.zshrc` or `.bash_profile`:

```bash
# Set up heads-down function with your preferred defaults
function heads-down() {
  SLACK_TOKEN="xoxp-your-token-here" /path/to/heads-down-slacker -duration 1h -channels "YOUR_CHANNEL_IDS"
}

function heads-up() {
  SLACK_TOKEN="xoxp-your-token-here" /path/to/heads-down-slacker -down=false
}
```