// Package keyboards containes all bot's keyboards.
package keyboards

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// BackToMain returns back to main keyboard.
func BackToMain() tgbotapi.InlineKeyboardMarkup {
	return backToMainRu
}

var backToMainRu = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("В главное меню", "backToMain"),
	),
)
