// Package bot handles bot activities.
package bot

import (
	"context"

	"github.com/Slava02/helperBot/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Message handles message requests.
func (h *handler) Message(ctx context.Context, tgbot *tgbotapi.BotAPI, update tgbotapi.Update) {
	deleteLastMessage(tgbot, update.Message.Chat.ID, update.Message.MessageID)

	user := &models.User{
		TelegramID: update.Message.From.ID,
		Username:   update.Message.From.UserName,
		FirstName:  update.Message.From.FirstName,
		LastName:   update.Message.From.LastName,
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	if UserEnterMode == true {
		var err error
		msg, err = h.message.Save(ctx, msg, user, mediaType)
		if err != nil {
			h.logger.Error(err)
		}
	} else {
		msg = h.message.Wrong(ctx, msg, user)
	}

	_, err := tgbot.Send(msg)
	if err != nil {
		h.logger.Error(err)
	}
}
