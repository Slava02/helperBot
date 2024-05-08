// package callback implements callback request functionalities.
package callback

import (
	"context"

	"cleanBot/models"
	"cleanBot/pkg/keyboards"
	"cleanBot/pkg/messages"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Help shows help message.
func (c *call) Help(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User) tgbotapi.MessageConfig {
	msg.Text = messages.Help(user)
	msg.ReplyMarkup = keyboards.BackToMain()

	return msg
}
