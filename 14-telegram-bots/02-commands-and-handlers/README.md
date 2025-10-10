# Commands and Handlers

## Overview
Learn how to handle specific bot commands and create structured command handlers. Commands are the primary way users interact with Telegram bots, and this section teaches you how to implement them effectively.

## What You'll Learn
- Creating command handlers for `/start`, `/help`, etc.
- Understanding command routing
- Extracting command arguments
- Implementing multiple command handlers
- Command best practices

## Telegram Bot Commands
Commands in Telegram start with `/` and can include parameters:
- `/start` - Traditional bot initialization command
- `/help` - Display help information
- `/echo <text>` - Commands with arguments
- `/settings` - Navigate to settings

## Code Example: Command Handler Bot

The `command_bot.go` file demonstrates:
- Multiple command handlers
- `/start` - Welcome message
- `/help` - Display available commands
- `/echo` - Echo back user arguments
- `/info` - Display user information
- Fallback handler for unrecognized commands

## Key Concepts

### Command Handler
Command handlers are filters that match specific commands:
```go
dispatcher.AddHandler(handlers.NewCommand("start", startHandler))
```

### Extracting Arguments
Commands can have arguments after the command:
```go
args := ctx.Args() // Gets text after command
```

### Command Scope
You can set which chats your commands are available in:
- Private chats only
- Groups only
- All chats

## Running the Bot

```bash
# Set your bot token
export TELEGRAM_BOT_TOKEN="your-token-from-botfather"

# Run the bot
go run command_bot.go
```

## Testing Commands
Try these commands with your bot:
1. `/start` - See the welcome message
2. `/help` - View all available commands
3. `/echo Hello World` - Echo back "Hello World"
4. `/info` - Get your user information
5. `random text` - See the fallback response

## Best Practices

### 1. Always Implement /start and /help
These are standard commands users expect.

### 2. Provide Clear Command Descriptions
Use BotFather to set command descriptions:
```
/setcommands
start - Start the bot
help - Show available commands
echo - Echo back your message
info - Show your user info
```

### 3. Handle Invalid Arguments
Always validate and handle missing or invalid arguments gracefully.

### 4. Use Consistent Naming
Keep command names lowercase and descriptive.

## Common Patterns

### Command with Required Arguments
```go
if len(ctx.Args()) == 0 {
    return errors.New("please provide text to echo")
}
```

### Admin-Only Commands
Check user permissions before executing certain commands.

## Next Steps
- Add more custom commands
- Implement command argument validation
- Try group-specific commands
- Move on to: Callbacks and Interactions

## Resources
- [Telegram Bot Commands](https://core.telegram.org/bots#commands)
- [BotFather Command Configuration](https://core.telegram.org/bots#6-botfather)
- [gotgbot Command Handlers](https://pkg.go.dev/github.com/PaulSonOfLars/gotgbot/v2/ext/handlers)
