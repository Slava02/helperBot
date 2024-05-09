package callback

import (
	"context"
	"fmt"
	"github.com/Slava02/helperBot/models"
	"github.com/Slava02/helperBot/pkg/keyboards"
	"github.com/Slava02/helperBot/pkg/messages"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *call) SaveOrRemove(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User) tgbotapi.MessageConfig {
	msg.Text = messages.SaveOrRemove(user)
	msg.ReplyMarkup = keyboards.SaveOrRemove()

	return msg
}

func (c *call) Save(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User, media models.Media) tgbotapi.MessageConfig {
	msg.Text = fmt.Sprintf("Cохранить %d", media)
	msg.ReplyMarkup = keyboards.BackToMain()

	return msg
}

func (c *call) Remove(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User, media models.Media) tgbotapi.MessageConfig {
	msg.Text = fmt.Sprintf("Удалить %d", media)
	msg.ReplyMarkup = keyboards.BackToMain()

	return msg
}
