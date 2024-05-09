package keyboards

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func Media() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Книги", "book"),
			tgbotapi.NewInlineKeyboardButtonData("Кино", "movie"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("В главное меню", "backToMain"),
		),
	)
}
