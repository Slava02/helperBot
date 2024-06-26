// Package bot handles bot activities.
package bot

import (
	"context"
	"time"

	"github.com/Slava02/helperBot/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Command handles command requests.
func (h *handler) Command(ctx context.Context, tgbot *tgbotapi.BotAPI, update tgbotapi.Update) {
	deleteLastMessage(tgbot, update.Message.Chat.ID, update.Message.MessageID)

	user := &models.User{
		TelegramID: update.Message.From.ID,
		Username:   update.Message.From.UserName,
		FirstName:  update.Message.From.FirstName,
		LastName:   update.Message.From.LastName,
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	if update.Message.Command() == "start" {
		if h.account.IsExist(ctx, user) {
			msg = h.command.Help(ctx, msg, user)
		} else {
			user.JoinedAt = time.Now()

			err := h.account.Create(ctx, user)
			if err != nil {
				h.logger.Error(err)
			}

			msg = h.command.Help(ctx, msg, user)
		}
	}

	_, err := tgbot.Send(msg)
	if err != nil {
		h.logger.Error(err)
	}
}
