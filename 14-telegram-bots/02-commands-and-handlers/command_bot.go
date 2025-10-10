package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func main() {
	// Get bot token from environment variable
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		panic("TELEGRAM_BOT_TOKEN environment variable is not set")
	}

	// Create bot instance
	bot, err := gotgbot.NewBot(token, nil)
	if err != nil {
		panic("failed to create new bot: " + err.Error())
	}

	// Create updater and dispatcher
	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
			fmt.Println("an error occurred while handling update:", err.Error())
			return ext.DispatcherActionNoop
		},
		MaxRoutines: ext.DefaultMaxRoutines,
	})
	updater := ext.NewUpdater(dispatcher, nil)

	// Add command handlers
	dispatcher.AddHandler(handlers.NewCommand("start", startHandler))
	dispatcher.AddHandler(handlers.NewCommand("help", helpHandler))
	dispatcher.AddHandler(handlers.NewCommand("echo", echoHandler))
	dispatcher.AddHandler(handlers.NewCommand("info", infoHandler))

	// Add a fallback handler for unrecognized messages
	dispatcher.AddHandler(handlers.NewMessage(nil, fallbackHandler))

	// Start receiving updates
	err = updater.StartPolling(bot, &ext.PollingOpts{
		DropPendingUpdates: true,
		GetUpdatesOpts: &gotgbot.GetUpdatesOpts{
			Timeout: 9,
			RequestOpts: &gotgbot.RequestOpts{
				Timeout: 10,
			},
		},
	})
	if err != nil {
		panic("failed to start polling: " + err.Error())
	}

	fmt.Printf("%s has been started...\n", bot.User.Username)

	// Idle, to keep updates coming in
	updater.Idle()
}

// startHandler handles the /start command
func startHandler(bot *gotgbot.Bot, ctx *ext.Context) error {
	message := "👋 Welcome to the Command Bot!\n\n" +
		"I can handle various commands. Use /help to see what I can do!"

	_, err := ctx.EffectiveChat.SendMessage(bot, message, nil)
	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}
	return nil
}

// helpHandler handles the /help command
func helpHandler(bot *gotgbot.Bot, ctx *ext.Context) error {
	message := "📚 Available Commands:\n\n" +
		"/start - Start the bot\n" +
		"/help - Show this help message\n" +
		"/echo <text> - Echo back your message\n" +
		"/info - Get your user information\n\n" +
		"Try sending any command to see it in action!"

	_, err := ctx.EffectiveChat.SendMessage(bot, message, nil)
	if err != nil {
		return fmt.Errorf("failed to send help message: %w", err)
	}
	return nil
}

// echoHandler handles the /echo command with arguments
func echoHandler(bot *gotgbot.Bot, ctx *ext.Context) error {
	// Get the arguments (text after the command)
	args := ctx.Args()

	var message string
	if len(args) == 0 {
		message = "❌ Please provide some text to echo!\n\nUsage: /echo <your text>"
	} else {
		// Join all arguments into a single string
		text := strings.Join(args, " ")
		message = fmt.Sprintf("🔊 You said: %s", text)
	}

	_, err := ctx.EffectiveChat.SendMessage(bot, message, nil)
	if err != nil {
		return fmt.Errorf("failed to send echo message: %w", err)
	}
	return nil
}

// infoHandler handles the /info command
func infoHandler(bot *gotgbot.Bot, ctx *ext.Context) error {
	user := ctx.EffectiveUser

	message := fmt.Sprintf(
		"👤 User Information:\n\n"+
			"ID: %d\n"+
			"First Name: %s\n"+
			"Username: @%s\n"+
			"Language: %s",
		user.Id,
		user.FirstName,
		user.Username,
		user.LanguageCode,
	)

	_, err := ctx.EffectiveChat.SendMessage(bot, message, nil)
	if err != nil {
		return fmt.Errorf("failed to send info message: %w", err)
	}
	return nil
}

// fallbackHandler handles all other messages
func fallbackHandler(bot *gotgbot.Bot, ctx *ext.Context) error {
	message := "🤔 I don't understand that command.\n\n" +
		"Use /help to see available commands!"

	_, err := ctx.EffectiveChat.SendMessage(bot, message, nil)
	if err != nil {
		return fmt.Errorf("failed to send fallback message: %w", err)
	}
	return nil
}
