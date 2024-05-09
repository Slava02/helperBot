package keyboards

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func SaveOrRemove() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Сохранить", "save"),
			tgbotapi.NewInlineKeyboardButtonData("Удалить", "remove"),
		),
	)
}
