# Getting Started with Telegram Bots

## Overview
This section introduces you to Telegram bot development with Go. You'll learn how to set up your environment, install the gotgbot library, and create your first simple echo bot.

## What You'll Learn
- Installing and setting up gotgbot
- Creating a basic bot that responds to messages
- Understanding the bot update loop
- Handling text messages
- Basic bot configuration

## Setup Instructions

### 1. Initialize a Go Module
```bash
go mod init telegram-bot-tutorial
```

### 2. Install gotgbot
```bash
go get github.com/PaulSonOfLars/gotgbot/v2
go get github.com/PaulSonOfLars/gotgbot/v2/ext
```

### 3. Set Up Your Bot Token
Never hardcode your bot token! Use environment variables:
```bash
export TELEGRAM_BOT_TOKEN="your-bot-token-here"
```

## Code Example: Echo Bot

The `echo_bot.go` file demonstrates:
- Initializing the bot with your token
- Setting up an updater for long polling
- Creating a message handler
- Responding to user messages
- Graceful error handling

## Running the Bot

```bash
# Set your bot token
export TELEGRAM_BOT_TOKEN="your-token-from-botfather"

# Run the bot
go run echo_bot.go
```

## Testing Your Bot
1. Open Telegram
2. Search for your bot by username
3. Start a chat with `/start`
4. Send any message
5. The bot should echo back your message!

## Common Issues

### "Unauthorized" Error
- Check if your bot token is correct
- Make sure the environment variable is set properly

### Bot Not Responding
- Verify the bot is running (check console output)
- Ensure you're messaging the correct bot
- Check for error messages in the console

## Next Steps
Once you have the echo bot working:
- Try modifying the response message
- Add logging to see what messages are received
- Experiment with different message types
- Move on to the next section: Commands and Handlers

## Resources
- [gotgbot Long Polling Example](https://github.com/PaulSonOfLars/gotgbot/tree/master/examples)
- [Telegram Bot API - Getting Updates](https://core.telegram.org/bots/api#getting-updates)
