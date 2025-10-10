package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/callbackquery"
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

	// Add command handler
	dispatcher.AddHandler(handlers.NewCommand("start", startHandler))
	dispatcher.AddHandler(handlers.NewCommand("menu", menuHandler))

	// Add callback handlers for different menu options
	dispatcher.AddHandler(handlers.NewCallback(callbackquery.Equal("menu_main"), mainMenuCallback))
	dispatcher.AddHandler(handlers.NewCallback(callbackquery.Prefix("category_"), categoryCallback))
	dispatcher.AddHandler(handlers.NewCallback(callbackquery.Prefix("item_"), itemCallback))

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
	message := "👋 Welcome to the Interactive Menu Bot!\n\n" +
		"Use /menu to open the interactive menu."

	_, err := ctx.EffectiveChat.SendMessage(bot, message, nil)
	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}
	return nil
}

// menuHandler handles the /menu command and displays the main menu
func menuHandler(bot *gotgbot.Bot, ctx *ext.Context) error {
	keyboard := gotgbot.InlineKeyboardMarkup{
		InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
			{
				{Text: "📚 Category 1", CallbackData: "category_1"},
				{Text: "🎮 Category 2", CallbackData: "category_2"},
			},
			{
				{Text: "🎨 Category 3", CallbackData: "category_3"},
			},
		},
	}

	_, err := ctx.EffectiveChat.SendMessage(
		bot,
		"📋 Main Menu\n\nPlease select a category:",
		&gotgbot.SendMessageOpts{
			ReplyMarkup: keyboard,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to send menu: %w", err)
	}
	return nil
}

// mainMenuCallback handles the "Back to Main Menu" button
func mainMenuCallback(bot *gotgbot.Bot, ctx *ext.Context) error {
	// Answer the callback query to remove loading state
	_, err := ctx.CallbackQuery.Answer(bot, &gotgbot.AnswerCallbackQueryOpts{
		Text: "Returning to main menu...",
	})
	if err != nil {
		return fmt.Errorf("failed to answer callback: %w", err)
	}

	// Update the message with the main menu
	keyboard := gotgbot.InlineKeyboardMarkup{
		InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
			{
				{Text: "📚 Category 1", CallbackData: "category_1"},
				{Text: "🎮 Category 2", CallbackData: "category_2"},
			},
			{
				{Text: "🎨 Category 3", CallbackData: "category_3"},
			},
		},
	}

	_, _, err = ctx.CallbackQuery.Message.EditText(
		bot,
		"📋 Main Menu\n\nPlease select a category:",
		&gotgbot.EditMessageTextOpts{
			ReplyMarkup: keyboard,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to edit message: %w", err)
	}
	return nil
}

// categoryCallback handles category selection
func categoryCallback(bot *gotgbot.Bot, ctx *ext.Context) error {
	// Extract category number from callback data
	callbackData := ctx.CallbackQuery.Data
	categoryNum := strings.TrimPrefix(callbackData, "category_")

	// Answer the callback query
	_, err := ctx.CallbackQuery.Answer(bot, &gotgbot.AnswerCallbackQueryOpts{
		Text: fmt.Sprintf("Opening Category %s...", categoryNum),
	})
	if err != nil {
		return fmt.Errorf("failed to answer callback: %w", err)
	}

	// Create keyboard with items for this category
	keyboard := gotgbot.InlineKeyboardMarkup{
		InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
			{
				{Text: fmt.Sprintf("Item %s-A", categoryNum), CallbackData: fmt.Sprintf("item_%s_a", categoryNum)},
			},
			{
				{Text: fmt.Sprintf("Item %s-B", categoryNum), CallbackData: fmt.Sprintf("item_%s_b", categoryNum)},
			},
			{
				{Text: fmt.Sprintf("Item %s-C", categoryNum), CallbackData: fmt.Sprintf("item_%s_c", categoryNum)},
			},
			{
				{Text: "🔙 Back to Main Menu", CallbackData: "menu_main"},
			},
		},
	}

	// Update the message
	messageText := fmt.Sprintf("📂 Category %s\n\nSelect an item:", categoryNum)
	_, _, err = ctx.CallbackQuery.Message.EditText(
		bot,
		messageText,
		&gotgbot.EditMessageTextOpts{
			ReplyMarkup: keyboard,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to edit message: %w", err)
	}
	return nil
}

// itemCallback handles item selection
func itemCallback(bot *gotgbot.Bot, ctx *ext.Context) error {
	// Extract item details from callback data
	callbackData := ctx.CallbackQuery.Data
	itemInfo := strings.TrimPrefix(callbackData, "item_")
	parts := strings.Split(itemInfo, "_")

	// Answer the callback query with a notification
	_, err := ctx.CallbackQuery.Answer(bot, &gotgbot.AnswerCallbackQueryOpts{
		Text: fmt.Sprintf("✅ You selected Item %s-%s", parts[0], strings.ToUpper(parts[1])),
	})
	if err != nil {
		return fmt.Errorf("failed to answer callback: %w", err)
	}

	// Create keyboard with back button
	keyboard := gotgbot.InlineKeyboardMarkup{
		InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
			{
				{Text: "🔙 Back to Category", CallbackData: fmt.Sprintf("category_%s", parts[0])},
			},
			{
				{Text: "🏠 Main Menu", CallbackData: "menu_main"},
			},
		},
	}

	// Update the message
	messageText := fmt.Sprintf(
		"✨ Item Details\n\n"+
			"Category: %s\n"+
			"Item: %s\n\n"+
			"This is a detailed view of the selected item.",
		parts[0], strings.ToUpper(parts[1]))

	_, _, err = ctx.CallbackQuery.Message.EditText(
		bot,
		messageText,
		&gotgbot.EditMessageTextOpts{
			ReplyMarkup: keyboard,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to edit message: %w", err)
	}
	return nil
}
