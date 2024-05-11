package keyboards

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func MediaOptions() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Добавить", "saveWhat"),
			tgbotapi.NewInlineKeyboardButtonData("Удалить", "removeWhat"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Рандом", "pickRandom"),
			tgbotapi.NewInlineKeyboardButtonData("Список", "list"),
		),
	)
}

func RemoveWhat(mediaObjets []string) tgbotapi.InlineKeyboardMarkup {
	var rows = make([][]tgbotapi.InlineKeyboardButton, 0)

	for _, m := range mediaObjets {
		row := tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(m, "r#"+m),
		)
		rows = append(rows, row)
	}

	backToMain := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("В главное меню", "backToMain"),
	)
	rows = append(rows, backToMain)

	return tgbotapi.NewInlineKeyboardMarkup(rows...)
}
