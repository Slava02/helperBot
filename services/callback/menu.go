// package callback implements callback request functionalities.
package callback

import (
	"context"

	"github.com/Slava02/helperBot/models"
	"github.com/Slava02/helperBot/pkg/keyboards"
	"github.com/Slava02/helperBot/pkg/messages"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Menu shows main menu.
func (c *call) Menu(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User) tgbotapi.MessageConfig {
	msg.Text = messages.Main(user)
	msg.ReplyMarkup = keyboards.Main()

	return msg
}
