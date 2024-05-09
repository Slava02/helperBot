// package account implements message request functionalities.
package message

import (
	"context"

	"github.com/Slava02/helperBot/models"
	"github.com/Slava02/helperBot/pkg/messages"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Wrong shows a wrong message.
func (m *msg) Wrong(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User) tgbotapi.MessageConfig {
	msg.Text = messages.Wrong(user)
	//msg.ReplyMarkup = keyboards.BackToMain()

	return msg
}
