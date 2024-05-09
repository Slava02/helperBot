package keyboards

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func MediaOptions() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Сохранить", "saveWhat"),
			tgbotapi.NewInlineKeyboardButtonData("Удалить", "removeWhat"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Случайный вариант", "pickRandom"),
		),
	)
}
