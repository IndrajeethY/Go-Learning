# Conversations

## Overview
This section demonstrates how to manage multi-step conversations with users using gotgbot's conversation handler. The conversation handler helps you track states across multiple messages, allowing you to guide users through complex interactions that require sequential steps.

## What You'll Learn
- Creating conversation handlers with state management
- Defining conversation entry points and states
- Managing multi-step user flows
- Handling conversation cancellation
- Using in-memory state storage
- Validating user input within conversations

## Code Example: Conversation Bot

The `main.go` file demonstrates a complete conversation bot that:
- Asks the user for their name
- Asks for their age with validation
- Handles state transitions between conversation steps
- Allows users to cancel the conversation at any point

### Key Components

#### 1. Conversation States
Define the different states in your conversation:
```go
const (
    NAME = "name"
    AGE  = "age"
)
```

#### 2. Conversation Handler Setup
Create a conversation handler with entry points, states, and exit handlers:
```go
dispatcher.AddHandler(handlers.NewConversation(
    []ext.Handler{handlers.NewCommand("start", start)},  // Entry points
    map[string][]ext.Handler{
        NAME: {handlers.NewMessage(noCommands, name)},    // State handlers
        AGE:  {handlers.NewMessage(noCommands, age)},
    },
    &handlers.ConversationOpts{
        Exits:        []ext.Handler{handlers.NewCommand("cancel", cancel)},
        StateStorage: conversation.NewInMemoryStorage(conversation.KeyStrategySenderAndChat),
        AllowReEntry: true,
    },
))
```

#### 3. State Management
- **Entry Handler**: The `/start` command initiates the conversation
- **State Handlers**: Each state has its own handler function
- **Exit Handler**: The `/cancel` command exits the conversation
- **State Storage**: In-memory storage tracks each user's conversation state

#### 4. State Transitions
Move between states using:
```go
return handlers.NextConversationState(NAME)  // Move to NAME state
return handlers.EndConversation()             // End the conversation
```

## Running the Bot

```bash
# Set your bot token
export TOKEN="your-token-from-botfather"

# Run the bot
go run main.go
```

## Testing the Conversation Flow

Try this conversation sequence:
1. Send `/start` - Bot asks for your name
2. Type your name (e.g., "Alice") - Bot greets you and asks for age
3. Type your age (e.g., "25") - Bot acknowledges and ends conversation

Or test the cancel flow:
1. Send `/start` - Bot asks for your name
2. Send `/cancel` - Bot cancels the conversation

Test input validation:
1. Send `/start` - Bot asks for your name
2. Type your name - Bot asks for age
3. Type "abc" (invalid) - Bot asks again for a valid number
4. Type "25" - Bot proceeds successfully

## Key Concepts

### Conversation Entry Points
Entry points are the handlers that start a conversation:
- Commands like `/start`
- Any handler that triggers the conversation flow

### State Handlers
Each conversation state maps to specific handlers:
- Only non-command text messages are accepted in this example
- Custom filters can be used for more complex scenarios

### State Storage
The conversation handler uses state storage to track each user's position:
- `KeyStrategySenderAndChat`: Separate conversations per user per chat
- In-memory storage: Simple but lost on restart
- Can be replaced with database storage for persistence

### Allow Re-Entry
When `AllowReEntry: true`, users can start a new conversation while already in one. Useful for:
- Letting users restart without canceling first
- Handling user mistakes gracefully

## Best Practices

### 1. Always Provide an Exit
Include a `/cancel` command or similar exit mechanism:
```go
Exits: []ext.Handler{handlers.NewCommand("cancel", cancel)},
```

### 2. Validate User Input
Always validate and handle invalid input:
```go
ageNumber, err := strconv.ParseInt(inputAge, 10, 64)
if err != nil {
    // Ask again with helpful message
    return handlers.NextConversationState(AGE)
}
```

### 3. Use HTML Escaping
Escape user input when sending back:
```go
html.EscapeString(inputName)
```

### 4. Clear Error Messages
Provide clear feedback when input is invalid:
```go
"This doesn't seem to be a number. Could you repeat?"
```

## Common Patterns

### Multi-Step Registration
```go
const (
    ASK_NAME  = "name"
    ASK_EMAIL = "email"
    ASK_PHONE = "phone"
    CONFIRM   = "confirm"
)
```

### Conditional States
Move to different states based on user input:
```go
if userResponse == "yes" {
    return handlers.NextConversationState(NEXT_STATE)
} else {
    return handlers.NextConversationState(ALTERNATIVE_STATE)
}
```

### Storing Conversation Data
Use context data to store information across states:
```go
ctx.Data["user_name"] = inputName
// Retrieve later:
userName := ctx.Data["user_name"].(string)
```

## Learning Path
Before diving into conversations, make sure you're comfortable with:
1. ✅ Basic message handling (01-getting-started)
2. ✅ Commands and handlers (02-commands-and-handlers)
3. ✅ Callbacks and interactions (03-callbacks-and-interactions)
4. ✅ Conversations (you are here)

## Next Steps
- Add more conversation states to the example
- Implement conversation data persistence
- Try conditional state transitions
- Build a survey or registration bot
- Explore conversation timeouts

## Resources
- [gotgbot Conversation Sample](https://github.com/PaulSonOfLars/gotgbot/tree/master/samples/conversationBot)
- [gotgbot Documentation](https://pkg.go.dev/github.com/PaulSonOfLars/gotgbot/v2)
- [Telegram Bot API](https://core.telegram.org/bots/api)
