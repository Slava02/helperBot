// Package keyboards containes all bot's keyboards.
package keyboards

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Main returns maun menu's keyboard.
func Main() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			//tgbotapi.NewInlineKeyboardButtonData("Ваша информация", "information"),
			tgbotapi.NewInlineKeyboardButtonData("Помощь", "help"),
			tgbotapi.NewInlineKeyboardButtonData("Медиатека", "media"),
		),
	)
}
