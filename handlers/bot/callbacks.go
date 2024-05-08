// Package bot handles bot activities.
package bot

import (
	"context"

	"cleanBot/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Callback handles callback requests.
func (h *handler) Callback(ctx context.Context, tgbot *tgbotapi.BotAPI, update tgbotapi.Update) {
	deleteLastMessage(tgbot, update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID)

	user := &models.User{
		TelegramID: update.CallbackQuery.From.ID,
		Username:   update.CallbackQuery.From.UserName,
		FirstName:  update.CallbackQuery.From.FirstName,
		LastName:   update.CallbackQuery.From.LastName,
	}

	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)

	var err error
	if err != nil {
		h.logger.Error(err)
	}

	switch update.CallbackQuery.Data {
	case "information":
		msg = h.callback.Information(ctx, msg, user)
	case "help":
		msg = h.callback.Help(ctx, msg, user)
	case "backToMain":
		msg = h.callback.Menu(ctx, msg, user)
	}

	_, err = tgbot.Send(msg)
	if err != nil {
		h.logger.Error(err)
	}
}
