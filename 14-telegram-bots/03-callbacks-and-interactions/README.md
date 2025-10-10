# Callbacks and Interactions

## Overview
Learn how to create interactive bots using inline keyboards and callback queries. This section covers building button-based interfaces that make your bot more engaging and user-friendly.

## What You'll Learn
- Creating inline keyboards with buttons
- Handling callback queries
- Updating messages dynamically
- Building interactive menus
- Callback data management

## Inline Keyboards vs Reply Keyboards

### Inline Keyboards
- Buttons appear below messages
- Send callback data to the bot
- Messages can be edited
- Better for quick actions

### Reply Keyboards
- Replace the user's keyboard
- Send text messages
- Good for navigation
- Covered in advanced topics

## Code Example: Interactive Menu Bot

The `callback_bot.go` file demonstrates:
- Creating inline keyboard buttons
- Handling button clicks (callbacks)
- Editing messages based on user selection
- Building a simple menu system
- Callback data parsing

## Key Concepts

### Creating Inline Keyboards
```go
keyboard := gotgbot.InlineKeyboardMarkup{
    InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
        {
            {Text: "Option 1", CallbackData: "option_1"},
            {Text: "Option 2", CallbackData: "option_2"},
        },
    },
}
```

### Handling Callbacks
```go
dispatcher.AddHandler(handlers.NewCallback(
    callbackquery.Prefix("option_"),
    handleCallback,
))
```

### Answering Callbacks
Always answer callback queries to remove loading state:
```go
ctx.CallbackQuery.Answer(bot, nil)
```

### Editing Messages
Update the message that contains the keyboard:
```go
ctx.CallbackQuery.Message.EditText(bot, newText, &gotgbot.EditMessageTextOpts{
    ReplyMarkup: newKeyboard,
})
```

## Running the Bot

```bash
# Set your bot token
export TELEGRAM_BOT_TOKEN="your-token-from-botfather"

# Run the bot
go run callback_bot.go
```

## Testing Interactions
1. Start the bot with `/start`
2. Click on the menu buttons
3. Navigate through different options
4. Observe how messages update
5. Try the "Back" buttons

## Best Practices

### 1. Always Answer Callbacks
```go
_, err := ctx.CallbackQuery.Answer(bot, nil)
```

### 2. Use Meaningful Callback Data
Structure your callback data for easy parsing:
- `menu_main` - Main menu
- `category_1` - Category selection
- `action_delete_123` - Action with ID

### 3. Handle Stale Callbacks
Messages might be edited or deleted, handle errors gracefully.

### 4. Limit Button Text Length
Keep button text short and clear (1-2 words when possible).

### 5. Provide Navigation
Always include "Back" or "Cancel" buttons in sub-menus.

## Common Patterns

### Confirmation Dialog
```go
keyboard := gotgbot.InlineKeyboardMarkup{
    InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
        {
            {Text: "✅ Confirm", CallbackData: "confirm_yes"},
            {Text: "❌ Cancel", CallbackData: "confirm_no"},
        },
    },
}
```

### Pagination
```go
{Text: "◀️ Prev", CallbackData: "page_1"},
{Text: "▶️ Next", CallbackData: "page_3"},
```

### Toggle Buttons
```go
{Text: "🔔 Notifications: ON", CallbackData: "toggle_notifications"}
```

## Advanced Topics
- URL buttons (open links)
- Switch inline query buttons
- Web app buttons
- Login buttons

## Next Steps
- Build a settings menu with multiple levels
- Implement pagination for long lists
- Add confirmation dialogs
- Move on to: Conversations

## Resources
- [Telegram Inline Keyboards](https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating)
- [Callback Query API](https://core.telegram.org/bots/api#callbackquery)
- [gotgbot Callback Handlers](https://pkg.go.dev/github.com/PaulSonOfLars/gotgbot/v2/ext/handlers)
