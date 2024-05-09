// package callback implements callback request functionalities.
package callback

import (
	"context"

	"github.com/Slava02/helperBot/models"
	"github.com/Slava02/helperBot/pkg/keyboards"
	"github.com/Slava02/helperBot/pkg/messages"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Information shows user's information.
func (c *call) Information(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User) tgbotapi.MessageConfig {
	msg.Text = messages.Information(user)
	msg.ReplyMarkup = keyboards.BackToMain()

	return msg
}
