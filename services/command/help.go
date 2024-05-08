package command

import (
	"cleanBot/models"
	"cleanBot/pkg/keyboards"
	"cleanBot/pkg/messages"
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *cmd) Help(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User) tgbotapi.MessageConfig {
	msg.Text = messages.Help(user)
	msg.ReplyMarkup = keyboards.BackToMain()

	return msg
}
